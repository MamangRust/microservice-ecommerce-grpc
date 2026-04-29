import re
import os

def parse_logs(file_path):
    with open(file_path, 'r') as f:
        lines = f.readlines()

    packages = {}
    current_package = None
    
    # Regex patterns
    pkg_start_re = re.compile(r'=== RUN\s+(\w+)')
    pkg_end_re = re.compile(r'--- (PASS|FAIL):\s+(\w+)')
    pkg_summary_re = re.compile(r'(ok|FAIL)\s+([\w\.\-/]+)\s+([\d\.]+s)')
    error_re = re.compile(r'^\s+([\w\.]+:\d+:.*)')
    panic_re = re.compile(r'panic:.*|test panicked:.*')

    i = 0
    while i < len(lines):
        line = lines[i].strip()
        
        # Detect package summary line
        pkg_summary_match = pkg_summary_re.search(line)
        if pkg_summary_match:
            status, pkg_name, duration = pkg_summary_match.groups()
            if pkg_name not in packages:
                packages[pkg_name] = {'status': status, 'tests': [], 'duration': duration}
            else:
                packages[pkg_name]['status'] = status
                packages[pkg_name]['duration'] = duration
            i += 1
            continue

        # Detect individual test pass/fail
        pkg_end_match = pkg_end_re.search(line)
        if pkg_end_match:
            status, test_name = pkg_end_match.groups()
            # Find the package this test belongs to (heuristic: look back for ok/FAIL lines is hard, 
            # so we'll just collect tests and associate them with the next package summary found or current)
            # Actually, Go logs usually show tests before the package summary.
            # Let's keep a list of tests and errors for the "current" context.
            i += 1
            continue

        i += 1

    # Second pass: More robust extraction for failures
    report = "# Test Results Summary\n\n"
    
    # Simple extraction: Just find all FAIL blocks
    failures = []
    current_failure = None
    
    for i, line in enumerate(lines):
        if "--- FAIL:" in line:
            if current_failure:
                failures.append(current_failure)
            current_failure = {'test': line.strip(), 'logs': []}
        elif "--- PASS:" in line:
            if current_failure:
                failures.append(current_failure)
                current_failure = None
        elif current_failure:
            if line.strip() and not line.startswith("==="):
                current_failure['logs'].append(line.strip())
        
        # Package summary
        pkg_m = pkg_summary_re.search(line)
        if pkg_m:
            if current_failure:
                failures.append(current_failure)
                current_failure = None

    if current_failure:
        failures.append(current_failure)

    # Generate Report
    report += "## Overall Context\n"
    report += f"- **Total Lines Analyzed**: {len(lines)}\n"
    report += f"- **Status**: See details below\n\n"

    report += "## Failed Packages\n"
    pkg_fails = set(re.findall(r'FAIL\s+([\w\.\-/]+)', "".join(lines)))
    for pf in sorted(pkg_fails):
        report += f"- `{pf}`\n"
    
    report += "\n## Detailed Failures\n"
    for f in failures:
        if not f['logs']: continue
        report += f"### {f['test']}\n"
        # Extract unique or important error messages
        important_logs = []
        for l in f['logs']:
            if any(x in l for x in ["Error:", "panic:", "Failed", "context deadline exceeded", "NOT_FOUND", "UNAUTHORIZED"]):
                important_logs.append(l)
        
        if important_logs:
            report += "```\n" + "\n".join(important_logs[:15]) + ("\n..." if len(important_logs) > 15 else "") + "\n```\n\n"
        else:
            report += "```\n" + "\n".join(f['logs'][:5]) + "\n```\n\n"

    return report

if __name__ == "__main__":
    src = "/home/hoover/Projects/golang/microservice-ecommerce-grpc/hello.txt"
    bak = src + ".bak"
    if not os.path.exists(bak):
        os.rename(src, bak)
        print(f"Backup created: {bak}")
    
    report_content = parse_logs(bak if os.path.exists(bak) else src)
    with open(src, 'w') as f:
        f.write(report_content)
    print(f"Summary written to {src}")
