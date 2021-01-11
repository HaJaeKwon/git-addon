# golang 으로 개발해보는 git addon

## 1. git user.name, user.email 설정 변경
* 미리 환경변수에 사용할 user.name, user.email 정보를 등록해두고 user 정보를 간단하게 스위치해주는 기능을 제공
* 환경변수 추가
  * GIT_NAME_{domain} / GIT_EMAIL_{domain} 환경변수를 추가
  * GIT_NAME_google=JaekwonHa
  * GIT_EMAIL_google=hazxz1116@gmail.com
* build 후 binary 경로 등록
  * make user
  * cp bin/gituser /usr/local/bin/gituser
  * gituser google

## 2. stash 수행 후에 checkout, pull
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

## 3. Merge Commit 에서 사용할 Commit log 생성기
* git log format pretty 자기 브런치꺼 까지만
