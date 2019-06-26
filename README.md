# gdeyamlOperator

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
|  latest_mode    |  push or build , choose one mode to identify latest tag to you    | push    |
|  pull-pattern    |  (pull)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}"    | null    |
|  push-pattern    |  (push)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}    | null    |
|  stage    |  replace stage , new stage content    | null    |
|  inputfile    |  input file name , such as deploy.yml    | null    |
|  ouputfile    |  output file name , such as deploy-out.yml    | tmp_out.yml    |
|  user   |  user for docker login    | null   |
|  password   |  password for docker login    | null   |
|  push   |  push this image , default is false    | false   |
|  version   |  prints current binary version    | false   |
|  list    |  After sort tag list , we only deal with these top'number tags    | 5    |
