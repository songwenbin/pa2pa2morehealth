import os

#company.txt 位置
company_path = 'companyall.txt'
#已生成excel的存放路径
xlsx_path = 'D:\\company\\firstToExcel\\'
#去重后txt存放路径
dedup_path = 'D:/company/txtToExcel/'

f = open(company_path,'r', encoding='UTF-8')
lines=f.readlines()
f.close()
companylist = set()

#读取公司txt 并add 至 companylist
for line in lines:
    companylist.add(''.join(line.split()))
    
dirs = os.listdir(xlsx_path)
for file in dirs:
    # 仅对xlsx文件进行处理
    if file.endswith(".xlsx"):
        #获取已生成公司 并从 companylist remove掉
        companylist.remove(file.split('.')[0]) 

#去重后公司名写入txt
new_file = open(dedup_path + 'dedupCompany1.txt','w')
new_file.write('\n'.join(companylist))
new_file.close()