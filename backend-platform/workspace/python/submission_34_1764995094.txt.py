
import sys
import time
import resource

input = sys.stdin.readline

def mem_kb():
    return resource.getrusage(resource.RUSAGE_SELF).ru_maxrss


__mem_before = mem_kb()
__start = time.perf_counter()

# Python Code
def solve():
    T = int(input())
    for _ in range(T):
        A, B = map(int, input().split())
        print(A - B)
solve()

__end = time.perf_counter()
__mem_after = mem_kb()

print(f"\nTIME_MS={(__end-__start)*1000:.3f}")
print(f"MEMORY_KB={__mem_after - __mem_before}")
