## krx

한국거래소(krx) 에서 다운로드 받은 시세데이터를 사용합니다.

### 데이터 활용
전종목 시세가 담긴 CSV 데이터를 다운로드 받으면 인코딩 문제로 바로 열람이 불가하고, UTF-8로 인코딩 변환이 필요합니다. 아래의 명령어를 활용합니다.
```shell
iconv -c -f euc-kr -t utf-8 ./downloaded.csv > data.csv
```

### 부록
데이터 접근 URL) http://data.krx.co.kr/contents/MDC/MDI/mdiLoader/index.cmd?menuId=MDC0201020101

사이트맵 path) 좌측 메뉴 - 주식 - 종목시세 - 전종목 시세


