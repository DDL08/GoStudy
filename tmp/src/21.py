def deduplicate_domains(input_file, output_file):
    # 打开输入文件读取域名
    with open(input_file, 'r', encoding='utf-8') as f:
        domains = f.readlines()
    
    # 去重
    unique_domains = set(domain.strip() for domain in domains)
    
    # 将去重后的域名写入输出文件
    with open(output_file, 'w', encoding='utf-8') as f:
        for domain in sorted(unique_domains):  # 按字母顺序排序
            f.write(domain + '\n')
    
    print(f"去重完成，共提取到 {len(unique_domains)} 个唯一域名，结果已保存到 {output_file}")

# 调用函数
input_file = r'C:\Users\LEGION\Desktop\2.txt'  # 输入文件路径
output_file = r'C:\Users\LEGION\Desktop\3.txt'  # 输出文件路径
deduplicate_domains(input_file, output_file)
#去重