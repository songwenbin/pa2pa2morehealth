import os
import json
import pandas as pd


class TxtToExcel:
    # 常量定义
    # 企业txt存放路径   
    path = '../first/'
    # excel存储路径
    savepath = '../firstToExcel/'
    # 正则去除html标签
#    dr = re.compile(r'<[^>]+>',re.S)
    
    def __init__(self):
        dirs = os.listdir( self.path )
        for file in dirs:
            # 仅对txt文件进行处理
            if file.endswith(".txt"):
                keyword = file.split('.')[0]
#                print(keyword)
                f = open(self.path + file,'r', encoding='UTF-8')
                lines = f.readlines()
                f.close()
                i = 1    #行号
                table_dict = {}
                for line in lines:
                    if(line[:1]=='{' and line[-2:-1] == '}') or line == '\n':
                        jsonstr = line
                        if line == '\n' and i > 1:
                            jsonstr = '{"recordsTotal":0}'
                        if line == '\n' and i == 1:
                            break
                        jsonobj = json.loads(jsonstr)
                        if i == 1:
                            table_dict['基本信息'] = self.get_base_info(jsonobj)
                        elif i == 2:
                            table_dict['股东信息'] = self.get_holder_info(jsonobj)
                        elif i == 3:
                            table_dict['主要人员'] = self.get_staff_info(jsonobj)
                        elif i == 4:
                            if int(jsonobj.get('recordsTotal')) > 0:
                                print(keyword+'=====4=====:'+str(jsonobj.get('recordsTotal')))
                        elif i == 5:
                            if jsonobj.get('recordsTotal') > 0:
                                print(keyword+'====5====:'+str(jsonobj.get('recordsTotal')))
                        elif i == 6:
                            table_dict['清算信息'] = self.get_liquidation_info(jsonobj)
                        elif i == 7:
                            table_dict['变更信息'] = self.get_alter_info(jsonobj)
                        elif i == 8:
                            table_dict['动产抵押登记信息'] = self.get_mortgage_info(jsonobj)
                        elif i == 9:
                            table_dict['股权出质登记信息'] = self.get_pledgor_info(jsonobj)
                        elif i == 10:
                            if jsonobj.get('recordsTotal') > 0:
                                print(keyword+'====10====:'+str(jsonobj.get('recordsTotal')))
                        elif i == 11:
                            if jsonobj.get('recordsTotal') > 0:
                                print(keyword+'====11====:'+str(jsonobj.get('recordsTotal')))
                        elif i == 12:
                            table_dict['司法协助信息'] = self.get_legal_info(jsonobj)
                        elif i == 13:
                            table_dict['行政许可信息'] = self.get_lic_info(jsonobj)
                        elif i == 14:
                            table_dict['行政处罚信息'] = self.get_penalties_info(jsonobj)
                        elif i == 15:
                            table_dict['列入经营异常名录信息'] = self.get_abnormal_info(jsonobj)
                        elif i == 16:
                            table_dict['列入严重违法失信企业名单(黑名单)信息'] = self.get_blacklist_info(jsonobj)
                        else:
                            break
                        i += 1
                    else:
                        continue
                # table_dict为空时，不生成excel
                if len(table_dict) == 0:
                    continue
                # 生成excel
                with pd.ExcelWriter(self.savepath + keyword+'.xlsx') as writer:
                    for sheet_name in table_dict:
                        table_dict[sheet_name].to_excel( writer, sheet_name = sheet_name, index = None)
                            
    # 第一行：基本信息 
    def get_base_info(self,jsonobj):
        # base_table 基础数据
        base_table = {}
        if 'result' in jsonobj :
            if jsonobj.get('result') != {}:
                result = jsonobj.get('result')
                base_table['机构名称'] = '' if result.get('entName') == None else result.get('entName')
                base_table['登记状态'] = '' if result.get('regState_CN') == None else result.get('regState_CN')
                base_table['注册号'] = '' if result.get('regNo') == None else result.get('regNo')
                base_table['统一社会信用代码'] = '' if result.get('uniscId') == None else result.get('uniscId')
                base_table['法人代表'] = '' if result.get('name') == None else result.get('name')
                base_table['住所'] = '' if result.get('dom') == None else result.get('dom')
                base_table['类型'] = '' if result.get('entType_CN') == None else result.get('entType_CN')
                base_table['成立日期'] = '' if result.get('estDate') == None else result.get('estDate')
                base_table['营业期限'] = '' if result.get('opFrom') == None else result.get('opFrom') +( '' if result.get('opTo') == None else (' 至 '+result.get('opTo')))
                base_table['核准日期'] = '' if result.get('apprDate') == None else result.get('apprDate')
                base_table['登记机关'] = '' if result.get('regOrg') == None else result.get('regOrg')
                base_table['注册资本'] = '' if result.get('regCap') == None else (str(result.get('regCap'))+ '万'+ ('人民币' if result.get('regCapCur_CN') == None else result.get('regCapCur_CN')))
                base_table['经营范围'] = '' if result.get('opScope') == None else result.get('opScope')
            else :
                base_table = ({'机构名称': '','登记状态':'','注册号':'','统一社会信用代码':'','法人代表': '','住所':'','类型':'','成立日期':'','营业期限': '','核准日期':'','登记机关':'','注册资本':'','经营范围':''})
        else:
            base_table = ({'机构名称': '','登记状态':'','注册号':'','统一社会信用代码':'','法人代表': '','住所':'','类型':'','成立日期':'','营业期限': '','核准日期':'','登记机关':'','注册资本':'','经营范围':''})
        return pd.DataFrame([base_table])
    
    # 第二行：股东信息
    def get_holder_info(self,jsonobj):
        holder_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                inv = '' if result.get('inv') == None else result.get('inv')
                liSubConAm = '' if result.get('liSubConAm') == None else result.get('liSubConAm')
                liAcConAm = '' if result.get('liAcConAm') == None else result.get('liAcConAm')
                holder_list.append({'股东名称': inv, '认缴出资额(万元)': liSubConAm,'实缴出资额(万元)':liAcConAm})
        else:
            holder_list.append({'股东名称': '', '认缴出资额(万元)': '','实缴出资额(万元)':''})
        holder_table = pd.DataFrame(holder_list, columns=['股东名称', '认缴出资额(万元)', '实缴出资额(万元)']) 
        return holder_table
    
    # 第三行：主要人员
    def get_staff_info(self,jsonobj):
        staff_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                position_CN = '' if result.get('position_CN') == None else result.get('position_CN')
                name = '' if result.get('name') == None else result.get('name')
                staff_list.append({'人员名称': name,'职位': position_CN })
        else:
            staff_list.append({'人员名称': '','职位': '' })
        staff_table = pd.DataFrame(staff_list, columns=['人员名称','职位' ])
        return staff_table
    
    
    # 第六行：清算信息
    def get_liquidation_info(self,jsonobj):
        liquidation_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                liqMem = '' if result.get('liqMem') == None else result.get('liqMem')
                liquidation_list.append({'清算组成员': liqMem})
        else:
            liquidation_list.append({'清算组成员': ''})
        liquidation_table = pd.DataFrame(liquidation_list, columns=['清算组成员'])
        return liquidation_table
    
    # 第七行：变更信息
    def get_alter_info(self,jsonobj):
        alter_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                altItem_CN = '' if result.get('altItem_CN') == None else result.get('altItem_CN')
                altDate = '' if result.get('altDate') == None else result.get('altDate')
                altBe = '' if result.get('altBe') == None else result.get('altBe')
                altAf = '' if result.get('altAf') == None else result.get('altAf')
                alter_list.append({'变更类型': altItem_CN,'变更时间':altDate,'变更前':altBe,'变更后':altAf})
        else:
            alter_list.append({'变更类型': '','变更时间':'','变更前':'','变更后':''})
        alter_table = pd.DataFrame(alter_list, columns=['变更类型','变更时间','变更前','变更后'])
        return alter_table
    
    # 第八行：动产抵押登记信息
    def get_mortgage_info(self,jsonobj):
        mortgage_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                morRegCNo = '' if result.get('morRegCNo') == None else result.get('morRegCNo')
                regiDate = '' if result.get('regiDate') == None else result.get('regiDate')
                regOrg_CN = '' if result.get('regOrg_CN') == None else result.get('regOrg_CN')
                priClaSecAm = '' if result.get('priClaSecAm') == None else (str(result.get('priClaSecAm')) + '万')
                publicDate = '' if result.get('publicDate') == None else result.get('publicDate')
                mortgage_list.append({'登记编号': morRegCNo,'登记日期':regiDate,'登记机关':regOrg_CN,'被担保债权数额':priClaSecAm,'公示日期':publicDate})
        else:
            mortgage_list.append({'登记编号': '','登记日期':'','登记机关':'','被担保债权数额':'','公示日期':''})
        mortgage_table = pd.DataFrame(mortgage_list, columns=['登记编号','登记日期','登记机关','被担保债权数额','公示日期'])
        return mortgage_table
    
    # 第九行：动产抵押登记信息
    def get_pledgor_info(self,jsonobj):
        pledgor_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                equityNo = '' if result.get('equityNo') == None else result.get('equityNo')
                equPleDate = '' if result.get('equPleDate') == None else result.get('equPleDate')
                pledgor = '' if result.get('pledgor') == None else result.get('pledgor')
                impAm = '' if result.get('impAm') == None else (str(result.get('impAm')) + '万元')
                impOrg = '' if result.get('impOrg') == None else result.get('impOrg')
                publicDate = '' if result.get('publicDate') == None else result.get('publicDate')
                pledgor_list.append({'登记编号': equityNo,'登记日期':equPleDate,'出质人':pledgor,'出质股权数额':impAm,'质权人':impOrg,'公示日期':publicDate})
        else:
            pledgor_list.append({'登记编号': '','登记日期':'','出质人':'','出质股权数额':'','质权人':'','公示日期':''})
        pledgor_table = pd.DataFrame(pledgor_list, columns=['登记编号','登记日期','出质人','出质股权数额','质权人','公示日期'])
        return pledgor_table
    
     # 第十二行：行政许可信息
    def get_legal_info(self,jsonobj):
        legal_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                executeNo = '' if result.get('executeNo') == None else result.get('executeNo')
                inv = '' if result.get('inv') == None else result.get('inv')
                froAuth = '' if result.get('froAuth') == None else result.get('froAuth')
                regCapCur = '' if result.get('regCapCur') == None else (str(result.get('regCapCur')) + '万' + '人民币' if result.get('regCapCur_CN') == None else result.get('regCapCur_CN'))
                frozState_CN = '' if result.get('frozState_CN') == None else result.get('frozState_CN')
                legal_list.append({'执行通知书文号': executeNo,'被执行人':inv,'股权数额':regCapCur,'执行法院':froAuth,'状态':frozState_CN})
        else:
            legal_list.append({'执行通知书文号': '','被执行人':'','股权数额':'','执行法院':'','状态':''})
        legal_table = pd.DataFrame(legal_list, columns=['执行通知书文号','股权数额','执行法院','出质股权数额','状态'])
        return legal_table
    
    # 第十三行：行政许可信息
    def get_lic_info(self,jsonobj):
        lic_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                licNo = '' if result.get('licNo') == None else result.get('licNo')
                licName_CN = '' if result.get('licName_CN') == None else result.get('licName_CN')
                licAnth = '' if result.get('licAnth') == None else result.get('licAnth')
                lic_list.append({'许可文件编号': licNo,'许可文件名称':licName_CN,'许可机关':licAnth})
        else:
            lic_list.append({'许可文件编号': '','许可文件名称':'','许可机关':''})
        lic_table = pd.DataFrame(lic_list, columns=['许可文件编号','许可文件名称','许可机关'])
        return lic_table
    
    # 第十四行：行政处罚信息
    def get_penalties_info(self,jsonobj):
        penalties_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                penDecNo = '' if result.get('penDecNo') == None else result.get('penDecNo')
                illegActType = '' if result.get('illegActType') == None else result.get('illegActType')
                penContent = '' if result.get('penContent') == None else result.get('penContent')
                penAuth_CN = '' if result.get('penAuth_CN') == None else result.get('penAuth_CN')
                penDecIssDate = '' if result.get('penDecIssDate') == None else result.get('penDecIssDate')
                penalties_list.append({'行政处罚决定书文号': penDecNo,'主要违法违规事实':illegActType,'行政处罚决定':penContent,'作出处罚决定机关':penAuth_CN,'作出决定日期':penDecIssDate})
        else:
            penalties_list.append({'行政处罚决定书文号': '','主要违法违规事实':'','行政处罚决定':'','作出处罚决定机关':'','作出决定日期':''})
        penalties_table = pd.DataFrame(penalties_list, columns=['行政处罚决定书文号','主要违法违规事实','行政处罚决定','作出处罚决定机关','作出决定日期'])
        return penalties_table
    
    # 第十五行：列入经营异常名录信息
    def get_abnormal_info(self,jsonobj):
        abnormal_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                speCause_CN = '' if result.get('speCause_CN') == None else result.get('speCause_CN')
                abntime = '' if result.get('abntime') == None else result.get('abntime')
                decOrg_CN = '' if result.get('decOrg_CN') == None else result.get('decOrg_CN')
                remExcpRes_CN = '' if result.get('remExcpRes_CN') == None else result.get('remExcpRes_CN')
                remDate = '' if result.get('remDate') == None else result.get('remDate')
                reDecOrg_CN = '' if result.get('reDecOrg_CN') == None else result.get('reDecOrg_CN')
                abnormal_list.append({'列入日期': abntime,'列入经营异常名录原因':speCause_CN,'作出决定机关(列入)':decOrg_CN,'移出日期':remDate,'移出经营异常名录原因':remExcpRes_CN,'作出决定机关(移出)':reDecOrg_CN})
        else:
            abnormal_list.append({'列入日期': '','列入经营异常名录原因':'','作出决定机关(列入)':'','移出日期':'','移出经营异常名录原因':'','作出决定机关(移出)':''})
        abnormal_table = pd.DataFrame(abnormal_list, columns=['列入日期','列入经营异常名录原因','作出决定机关(列入)','移出日期','移出经营异常名录原因','作出决定机关(移出)'])
        return abnormal_table
    
    # 第十六行：列入严重违法失信企业名单
    def get_blacklist_info(self,jsonobj):
        blacklist_list = []
        if int(jsonobj.get('recordsTotal')) > 0:
            for i in range(len(jsonobj.get('data'))):
                result = jsonobj.get('data')[i]
                serILLRea_CN = '' if result.get('serILLRea_CN') == None else result.get('serILLRea_CN')
                abntime = '' if result.get('abntime') == None else result.get('abntime')
                decOrg_CN = '' if result.get('decOrg_CN') == None else result.get('decOrg_CN')
                remExcpRes_CN = '' if result.get('remExcpRes_CN') == None else result.get('remExcpRes_CN')
                remDate = '' if result.get('remDate') == None else result.get('remDate')
                reDecOrg_CN = '' if result.get('reDecOrg_CN') == None else result.get('reDecOrg_CN')
                blacklist_list.append({'列入日期': abntime,'列入严重违法失信企业名单(黑名单)原因':serILLRea_CN,'作出决定机关(列入)':decOrg_CN,'移出日期':remDate,'移出经营异常名录原因':remExcpRes_CN,'作出决定机关(移出)':reDecOrg_CN})
        else:
            blacklist_list.append({'列入日期': '','列入严重违法失信企业名单(黑名单)原因':'','作出决定机关(列入)':'','移出日期':'','移出经营异常名录原因':'','作出决定机关(移出)':''})
        blacklist_table = pd.DataFrame(blacklist_list, columns=['列入日期','列入严重违法失信企业名单(黑名单)原因','作出决定机关(列入)','移出日期','移出经营异常名录原因','作出决定机关(移出)'])
        return blacklist_table
    
TxtToExcel()