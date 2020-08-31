# golang 으로 개발해보는 git command helper

### stash 수행 후에 checkout, pull

* init
  * only stash, pull
* init -d
  * develop branch로 이동
* init -b feature/issue-1
  * 특정 branch로 이동

init -> stash, checkout, pull

1. git 명령어 설치되어 있는지 확인
2. .git 폴더가 있는지 확인
3. stash 수행
4. checkout 수행
5. pull 수행

### Merge Commit 에서 사용할 Commit log 생성기

* git log format pretty 자기 브런치꺼 까지만

### checkout 단축키, 없으면 자동으로 -b

co add-schdule -f
co add-schdule -h
co add-schedule

feature/
hotfix/
