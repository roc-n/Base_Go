n, k, a, b = map(int, input().split())

ab = list(map(int, input().split()))  # 能力值
co = list(map(int, input().split()))  # 合作值

# 统计当前滑动窗口内符合条件的同学数量
cnt = 0
for i in range(k - 1):
    if ab[i] >= a and co[i] >= b:
        cnt += 1

result = 0
for i in range(k - 1, n):
    if ab[i] >= a and co[i] >= b:
        cnt += 1
    if i >= k:
        if ab[i - k] >= a and co[i - k] >= b:
            cnt -= 1
    if cnt == k:
        result += 1
print(result)