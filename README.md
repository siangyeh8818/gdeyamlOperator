# gdeyamlOperator

## 主要操作

- gettag : 返回該image 最新的tag並儲存到"getImageLatestTag_result.txt"這個檔案
- snapshot : 將k8s環境上的deploy,statefulset,daemonset,cronjob 輸出成gdeyaml格式
- nexus_api : 對Nexus repository 的API操作
- promote : 對Nexus repository 的資產進行搬移或是複製類型的搬移, 支援將gdeyaml格式的image 進行promote
- gitclone : 對於git clone的動作 , 參數以branch優先 , 若branch不存在會找tag 做clone
- git : git 相關操作 , 類似clone,checkoot,commit 等等
- replace : 將environment格式的replace內容取代gdeyaml的指定內容
- new-release : 為了開新的branch (gdeyaml & base的repo) , 並把new branch打上去gdeyaml文件
- imagedump : dump出k8s上的image 並可支援push

## 使用前須安裝

須能使用jq指令和docker指令

```sh
# CentOS
yum -y install jq
```

## 用法

./getImageLatestTag --imagename dockerhub.pentium.network/grafana

主要動作的flag

|  flag  |                                                               description                                                                | default value |
| :----: | :--------------------------------------------------------------------------------------------------------------------------------------: | :-----------: |
| action | action", "gettag", "choose 'gettag' or 'snapshot' or 'promote' or 'gitclone' or 'replace' or 'imagedump' or 'nexus_api' or 'new-release' |    gettag     |

Git相關的flag

|      flag      |                     description                      | default value |
| :------------: | :--------------------------------------------------: | :-----------: |
|    git-url     |                   url for git repo                   |     null      |
|   clone-path   |              folder path for git clone               |     null      |
| git-repo-path  |                directory for git-repo                |     null      |
|    git-user    |                  user for git clone                  |     null      |
|   git-token    |                 token for git clone                  |     null      |
|   git-branch   |                 branch for git repo                  |     null      |
| git-new-branch | New branch for git repo, this branch will be created |     null      |
|    git-tag     |                   Tag for git repo                   |     null      |
|   git-action   |   git related operation , such as 'branch','push'    |     null      |

Docker相關的flag

| flag         |                                              description                                               | default value |
| ------------ | :----------------------------------------------------------------------------------------------------: | :-----------: |
| docker-login |                                   DockerHub url/IP for docekr login                                    |     null      |
| push         |                                            push this image                                             |     false     |
| push-pattern |            (push)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}            |               |
| pull-pattern |            (pull)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}            |               |
| imagename    | docker image , such as dockerhub.pentium.network/grafana (default "dockerhub.pentium.network/grafana") |               |
| list         |                     After sort tag list , we only deal with these top'number tags                      |       5       |
| latest-mode  |                     push or build , choose one mode to identify latest tag to you                      |     push      |

Nexus相關的flag

| flag                 |                                 說明                                  | default value |
| -------------------- | :-------------------------------------------------------------------: | :-----------: |
| nexus-api-method     | Http method for NexusAPI Request, such as 'GET','POST','PUT','DELETE' |               |
| nexus-req-body       |                   Requets body for NexusAPI Request                   |               |
| nexus-output-pattern |              Pattern for output by requesting Nexus-API               |               |
| promote-type         |                   Different model , 'move' or 'cp'                    |     move      |
| promote-destination  |                    Destination for repository name                    |               |
| promote-url          |           destination for you promoting image url (nexus)'            |               |
| promote-source       |     sourece(Repository name) for you promoting image url (nexus)'     |               |
