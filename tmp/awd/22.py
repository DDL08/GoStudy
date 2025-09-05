import requests

# 打开文件写入可访问的主机
with open(r'F:\addl\awd\hosts.txt', "w") as f:
    for i in range(1, 255):
        url = f"http://192-168-1-{i}.pvp6070.bugku.cn"
        try:
            resp = requests.get(url, timeout=0.1)
            if resp.status_code == 200:
                print(f"[+] 200 OK: {url}")
                f.write(url + "\n")
            else:
                print(f"[-] {resp.status_code}: {url}")
        except requests.RequestException:
            print(f"[x] Error accessing: {url}")
