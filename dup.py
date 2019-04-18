import os

#txt 路径
txt_path = './first'
#已生成excel的存放路径
xlsx_path = './firstToExcel'
#去重后txt存放路径
dedup_path = './txtToExcel/'

companylist = set()
#获取txt
readtxt = os.listdir(txt_path)
for txt in readtxt:
    # 仅对txt文件进行处理
    if txt.endswith(".txt"):
        companylist.add(txt.split('.')[0]) 
        
#读取excel
readexcel = os.listdir(xlsx_path)
for xlsx in readexcel:
    # 仅对txt文件进行处理
    if xlsx.endswith(".xlsx"):
        companylist.remove(xlsx.split('.')[0]) 

#去重后公司名写入txt
new_file = open(dedup_path + 'dedupCompany2.txt','w')
new_file.write('\n'.join(companylist))
new_file.close()