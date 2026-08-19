import json
import re
import sys

def parse_markdown_to_jsons(filename):
    with open(filename, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Find all JSON blocks
    # Looking for ### results\set1\... followed by ```json ... ```
    pattern = r'### results\\set1\\[^\n]+\n+```json\n(.*?)\n```'
    matches = re.findall(pattern, content, re.DOTALL)
    
    results = []
    for match in matches:
        try:
            data = json.loads(match)
            results.append(data)
        except Exception as e:
            print(f"Error parsing json: {e}")
    return results

def main():
    filename = 'all_benchmark_results.md'
    data = parse_markdown_to_jsons(filename)
    
    if not data:
        print("No data found!")
        return
        
    print("=== Throughput vs. Skew ===")
    print("| Skew (Theta) | OCC (TPS) | PESSIMISTIC (TPS) | SSI (TPS) |")
    print("|---|---|---|---|")
    skews = [0.0, 0.4, 0.6, 0.8, 1.0, 1.2]
    for s in skews:
        occ = next((d['throughput_tps'] for d in data if d['strategy'] == 'OCC' and d['concurrency'] == 50 and float(d['skew_theta']) == s), None)
        pess = next((d['throughput_tps'] for d in data if d['strategy'] == 'PESSIMISTIC' and d['concurrency'] == 50 and float(d['skew_theta']) == s), None)
        ssi = next((d['throughput_tps'] for d in data if d['strategy'] == 'SSI' and d['concurrency'] == 50 and float(d['skew_theta']) == s), None)
        print(f"| {s} | {occ:.1f} | {pess:.1f} | {ssi:.1f} |")
        
    print("\n=== Throughput vs. Concurrency ===")
    print("| Concurrency | OCC (TPS) | PESSIMISTIC (TPS) | SSI (TPS) |")
    print("|---|---|---|---|")
    concs = [10, 50, 100, 250]
    for c in concs:
        occ = next((d['throughput_tps'] for d in data if d['strategy'] == 'OCC' and float(d['skew_theta']) == 0.8 and d['concurrency'] == c), None)
        pess = next((d['throughput_tps'] for d in data if d['strategy'] == 'PESSIMISTIC' and float(d['skew_theta']) == 0.8 and d['concurrency'] == c), None)
        ssi = next((d['throughput_tps'] for d in data if d['strategy'] == 'SSI' and float(d['skew_theta']) == 0.8 and d['concurrency'] == c), None)
        print(f"| {c} | {occ:.1f} | {pess:.1f} | {ssi:.1f} |")
        
    print("\n=== Latency C=250, Theta=0.8 ===")
    for strat in ['OCC', 'PESSIMISTIC', 'SSI']:
        d = next((d for d in data if d['strategy'] == strat and float(d['skew_theta']) == 0.8 and d['concurrency'] == 250), None)
        if d:
            print(f"Strategy: {strat}")
            print(f"  app_latency: P95={d['app_latency']['p95']/1000000:.2f}ms, P99={d['app_latency']['p99']/1000000:.2f}ms")
            print(f"  db_latency: P95={d['db_latency']['p95']/1000000:.2f}ms, P99={d['db_latency']['p99']/1000000:.2f}ms")
            print(f"  cc_wait_latency: P95={d['cc_wait_latency']['p95']/1000000:.2f}ms, P99={d['cc_wait_latency']['p99']/1000000:.2f}ms")
            print(f"  wal_wait_proxy: P95={d['wal_wait_proxy']['p95']/1000000:.2f}ms, P99={d['wal_wait_proxy']['p99']/1000000:.2f}ms")
            print(f"  client_e2e_latency: P95={d['client_e2e_latency']['p95']/1000000:.2f}ms, P99={d['client_e2e_latency']['p99']/1000000:.2f}ms")
            
    print("\n=== Dead Tuples ===")
    for d in data:
        print(f"{d['strategy']} c={d['concurrency']} s={d['skew_theta']}: {d.get('dead_tuples_per_commit', 'N/A')}")

if __name__ == '__main__':
    main()
