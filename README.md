# ไมโครเซอร์วิส

แอปพลิเคชันไมโครเซอร์วิสที่พัฒนาด้วยภาษา Go รองรับการเชื่อมต่อหลายฐานข้อมูลและ RESTful API

## คุณสมบัติ

- 🗄️ **รองรับหลายฐานข้อมูล**: PostgreSQL, MongoDB, MySQL, MSSQL
- 🚀 **เว็บเฟรมเวิร์ค**: Gin HTTP framework
- 📝 **เอกสาร API**: Swagger/OpenAPI integration
- 🔐 **การยืนยันตัวตน**: JWT Authentication
- 📊 **จัดการไฟล์ Excel**: ประมวลผลไฟล์ Excel ด้วย excelize
- 🌐 **เว็บออโตเมชัน**: Chrome DevTools Protocol support
- 📋 **การบันทึกข้อมูล**: Zap logging library
- 🔧 **การตั้งค่าสภาพแวดล้อม**: dotenv support

## เทคโนโลยีที่ใช้

- **Go 1.23.3**
- **Gin** - HTTP Web Framework
- **ฐานข้อมูล**:
  - PostgreSQL (lib/pq)
  - MongoDB (mongo-driver)
  - MySQL (go-sql-driver/mysql)
  - MSSQL (go-mssqldb)
- **การยืนยันตัวตน**: golang-jwt/jwt
- **เอกสาร**: Swagger (swaggo)
- **การบันทึกข้อมูล**: uber-go/zap
- **การประมวลผลไฟล์**: excelize/v2

## สิ่งที่ต้องเตรียม

- Go 1.23.3 ขึ้นไป
- Docker (สำหรับฐานข้อมูลในคอนเทนเนอร์)
- การเข้าถึงฐานข้อมูลทั้งหมดที่กำหนดค่าไว้

## การติดตั้ง

1. โคลน repository:
```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
cd YOUR_REPO_NAME
```

2. ติดตั้ง dependencies:
```bash
go mod download
```

3. คัดลอกไฟล์ environment:
```bash
cp .env.example .env
```

4. กำหนดค่าการเชื่อมต่อฐานข้อมูลใน `.env`

## การตั้งค่า

แอปพลิเคชันใช้ตัวแปรสภาพแวดล้อมสำหรับการกำหนดค่า ตัวแปรสำคัญได้แก่:

- สตริงการเชื่อมต่อฐานข้อมูลสำหรับ PostgreSQL, MongoDB, MySQL, MSSQL
- คีย์ลับ JWT
- พอร์ตเซิร์ฟเวอร์ (ค่าเริ่มต้น: 8080)
- ระดับและพาธของ log

## การรันแอปพลิเคชัน

### สำหรับพัฒนา

```bash
# รันโดยตรง
go run main.go

# หรือใช้ air สำหรับ hot reload
air
```

### สำหรับ Production

```bash
# Build
go build -o microservice main.go

# Run
./microservice
```

### ใช้ Docker

```bash
# Build image
docker build -t microservice .

# Run container
docker run -p 8080:8080 microservice
```

## เอกสาร API

เมื่อเซิร์ฟเวอร์ทำงาน ให้เข้าไปที่:
- Swagger UI: `http://localhost:8080/swagger/index.html`

## โครงสร้างโปรเจค

```
.
├── main.go              # จุดเริ่มต้นของแอปพลิเคชัน
├── go.mod               # ไฟล์ Go module
├── go.sum               # การตรวจสอบ dependencies
├── .env                 # ตัวแปรสภาพแวดล้อม
├── .gitignore           # กฎการ ignore ของ Git
├── databases/           # โมดูลการเชื่อมต่อฐานข้อมูล
│   ├── mongodb/
│   ├── mssql/
│   ├── mysql/
│   └── postgresql/
├── helper/              # ฟังก์ชัน utility
├── internals/           # ตรรกะภายในแอปพลิเคชัน
├── logs/                # การตั้งค่าการบันทึกข้อมูล
├── responses/           # โครงสร้าง response ของ API
├── routes/              # การกำหนด HTTP routes
├── storages/            # ที่เก็บไฟล์ (gitignored)
├── templates/           # ไฟล์เทมเพลต
└── docker-compose/      # การตั้งค่า Docker
```

## การตั้งค่าฐานข้อมูล

แอปพลิเคชันเชื่อมต่อกับหลายฐานข้อมูล ตรวจสอบให้แน่ใจว่าฐานข้อมูลทั้งหมดทำงานและเข้าถึงได้:

### ใช้ Docker Compose

```bash
docker-compose -f docker-compose/docker-compose.yml up -d
```

## ตัวแปรสภาพแวดล้อม

สร้างไฟล์ `.env` ด้วยตัวแปรต่อไปนี้:

```env
# PostgreSQL
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database

# MongoDB
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=your_database

# MySQL
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=your_user
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=your_database

# MSSQL
MSSQL_HOST=localhost
MSSQL_PORT=1433
MSSQL_USER=your_user
MSSQL_PASSWORD=your_password
MSSQL_DATABASE=your_database

# JWT
JWT_SECRET=your_jwt_secret_key

# Server
SERVER_PORT=8080
```

## การมีส่วนร่วมในการพัฒนา

1. Fork repository
2. สร้าง branch ใหม่ (`git checkout -b feature/amazing-feature`)
3. Commit การเปลี่ยนแปลง (`git commit -m 'Add some amazing feature'`)
4. Push ไปยัง branch (`git push origin feature/amazing-feature`)
5. เปิด Pull Request

## ใบอนุญาต

โปรเจคนี้ใช้ใบอนุญาต MIT License - ดูรายละเอียดในไฟล์ LICENSE

## การสนับสนุน

สำหรับการสนับสนุนและคำถาม กรุณาเปิด issue ใน GitHub repository
