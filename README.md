# SecondChance-Global-backend

> 🐕 유기견 입양 정보를 제공하는 글로벌 백엔드 API 서버

SecondChance-Global-backend는 미국 Petfinder API와 경기도 유기동물 보호 API를 통합하여 유기견 정보를 제공하는 REST API 서버입니다. 사용자는 미국과 한국의 유기견 정보를 조회하고 랜덤으로 추천받을 수 있습니다.

## 🚀 주요 기능

- **다국가 유기견 정보 조회**: 미국(Petfinder)과 한국(경기도) 유기견 데이터 통합
- **랜덤 유기견 추천**: 두 국가의 데이터를 랜덤하게 섞어서 제공
- **RESTful API**: 표준 HTTP 메서드를 사용한 깔끔한 API 설계
- **JWT 토큰 인증**: Petfinder API 접근을 위한 자동 토큰 관리
- **Docker 지원**: 컨테이너화된 배포 환경 제공

## 📋 API 엔드포인트

### Health Check
```
GET /api/v1/health
```
서버 상태 확인

### 유기견 조회
```
GET /api/v1/dogs/random
```
미국과 한국의 유기견 데이터를 랜덤하게 섞어서 반환

```
GET /api/v1/dogs/{id}?country={country}
```
특정 ID의 유기견 정보 조회
- `country`: "American" 또는 "Korean"

## 🏗️ 프로젝트 구조

```
SecondChance-Global-backend/
├── cmd/
│   └── main.go                 # 애플리케이션 진입점
├── internal/
│   ├── api/                    # 외부 API 클라이언트
│   │   ├── const.go           # API 상수 정의
│   │   ├── gyeonggi.go        # 경기도 API 클라이언트
│   │   └── petfinder.go       # Petfinder API 클라이언트
│   ├── config/                 # 설정 관리
│   │   └── config.go
│   ├── handler/                # HTTP 핸들러
│   │   ├── dog.go             # 유기견 관련 핸들러
│   │   └── health.go          # 헬스체크 핸들러
│   ├── middleware/             # 미들웨어
│   │   └── token.go           # Petfinder 토큰 미들웨어
│   ├── model/                  # 데이터 모델
│   │   ├── dog.go             # 유기견 모델
│   │   ├── gyeonggi.go        # 경기도 API 모델
│   │   └── petfinder.go       # Petfinder API 모델
│   ├── router/                 # 라우팅 설정
│   │   └── router.go
│   └── service/                # 비즈니스 로직
│       ├── dog.go             # 유기견 서비스
│       └── health.go          # 헬스체크 서비스
├── Dockerfile                  # Docker 설정
├── go.mod                      # Go 모듈 정의
└── README.md
```

## 🛠️ 기술 스택

- **언어**: Go 1.21
- **웹 프레임워크**: Chi Router
- **로깅**: Zerolog
- **설정 관리**: envconfig
- **환경 변수**: godotenv
- **CORS**: go-chi/cors
- **컨테이너**: Docker

## ⚙️ 설치 및 실행

### 1. 환경 변수 설정

`.env` 파일을 생성하고 다음 환경 변수를 설정하세요:

```env
# 서버 설정
SERVER_HOST=localhost
SERVER_PORT=8080

# Petfinder API 설정
PETFINDER_CLIENT_ID=your_client_id
PETFINDER_CLIENT_SECRET=your_client_secret
PETFINDER_ACCESS_TOKEN=your_access_token

# 경기도 API 설정
GYEONGGI_API_KEY=your_gyeonggi_api_key
```

### 2. 로컬 실행

```bash
# 의존성 설치
go mod download

# 애플리케이션 실행
go run cmd/main.go
```

### 3. Docker 실행

```bash
# Docker 이미지 빌드
docker build -t secondchance-global .

# 컨테이너 실행
docker run -p 8080:8080 --env-file .env secondchance-global
```

## 📊 데이터 소스

### 미국 데이터 (Petfinder)
- **API**: Petfinder API v2
- **데이터 타입**: 개(dog)
- **상태**: 입양 가능(adoptable)
- **총 페이지**: 1,715페이지
- **페이지당**: 10개 데이터

### 한국 데이터 (경기도)
- **API**: 경기도 공공데이터포털 유기동물보호 API
- **상태**: 보호중
- **총 데이터**: 2,129개
- **페이지당**: 10개 데이터

## 🔧 개발 가이드

### 코드 구조
- **Handler**: HTTP 요청/응답 처리
- **Service**: 비즈니스 로직 구현
- **API**: 외부 API 클라이언트
- **Model**: 데이터 구조 정의
- **Middleware**: 공통 기능 (인증, 로깅 등)

### 새로운 기능 추가
1. `internal/model/`에 데이터 모델 정의
2. `internal/service/`에 비즈니스 로직 구현
3. `internal/handler/`에 HTTP 핸들러 작성
4. `internal/router/`에 라우트 등록

## 📝 API 응답 예시

### 랜덤 유기견 조회
```json
{
  "message": "Random dogs selected from American and Korean data",
  "data": [
    {
      "id": 12345,
      "name": "Buddy",
      "age": "young",
      "images": ["url1", "url2"],
      "gender": "male",
      "breed": "Labrador",
      "location": {
        "country": "US",
        "city": "New York"
      },
      "shelter": {
        "name": "Petfinder",
        "contact": "555-1234",
        "email": "shelter@example.com"
      },
      "countryType": "American"
    }
  ]
}
```

### 헬스체크
```json
{
  "status": "ok",
  "message": "Server is running",
  "timestamp": "2024-01-01T00:00:00Z"
}
```
---
**SecondChance-Global-backend** - 유기견들에게 두 번째 기회를 제공하는 글로벌 플랫폼 🌍🐕
