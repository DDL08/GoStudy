import re

def extract_subdomains(input_file, output_file):
    # 打开输入文件读取URL
    with open(input_file, 'r', encoding='utf-8') as f:
        urls = f.readlines()
    
    # 提取子域名
    subdomains = set()  # 使用集合去重
    for url in urls:
        url = url.strip()  # 去掉多余的空格和换行符
        match = re.search(r'https?://[^/]+/([^/]+)', url)
        if match:
            subdomain = match.group(1)
            subdomains.add(subdomain)
    
    # 将子域名写入输出文件
    with open(output_file, 'w', encoding='utf-8') as f:
        for subdomain in sorted(subdomains):  # 按字母顺序排序
            f.write(subdomain + '\n')
    
    print(f"提取完成，共提取到 {len(subdomains)} 个子域名，结果已保存到 {output_file}")

# 调用函数
input_file = r'C:\Users\LEGION\Desktop\1.txt'  # 输入文件路径
output_file = r'C:\Users\LEGION\Desktop\2.txt'  # 输出文件路径
extract_subdomains(input_file, output_file)