import requests

alive = []

def check_alive(ip):
    target = f"http://192.168.1.{ip}.pvp6109.bugku.cn"
    try:
        response = requests.get(url=target, timeout=2)
        if response.status_code < 500:
            alive.append(target)
            print(target)
    except:
        pass

# 依次扫描 IP（无多线程）
for i in range(1, 255):
    check_alive(i)

# 写入结果文件
with open(r"F:\addl\awd\hosts.txt", 'w') as f:
    for host in alive:
        f.write(host + "\n")
