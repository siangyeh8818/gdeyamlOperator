# getImageLatestTag

返回該image 最新的tag並儲存到"getImageLatestTag_result.txt"這個檔案

# 使用前須安裝
須能使用jq指令和docker指令 <br>
Centos:<br>
  yum -y install jq

# 用法
./getImageLatestTag --imagename dockerhub.pentium.network/grafana

| flag      | 說明    | 預設值     |
| ---------- | :-----------:  | :-----------: |
|  imagename    | docker image , such as dockerhub.pentium.network/grafana     | dockerhub.pentium.network/grafana    |
| ---------- | :-----------:  | :-----------: |
|  list    |  After sort tag list , we only deal with these top'number tags    | 5    |
