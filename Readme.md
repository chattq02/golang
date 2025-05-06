router -> controller -> service -> repository -> database
// controller --> service --> repo --> models -> dbs
// wire: gen ra các khớp nối
file model được gen ra khi chạy do có genDTO ở mysql.go
fmt.Printf("user: %+v\n", user) -- giống với console.log
"sqlc generate" gen ra các func ở trong internal/database
////////////////////////////////////

Goose: Quản lý migration, giúp điều chỉnh schema cơ sở dữ liệu theo thời gian.
sqlc: Tạo mã Go an toàn về kiểu từ các câu lệnh SQL, giảm công sức viết các truy vấn thủ công.

1. Docker để tạo container:
   docker run --name mysqlgo8 -e MYSQL_ROOT_PASSWORD=root1234 -e MYSQL_DATABASE=shopdevgo -p 33306:3306 -d mysql:8

2. Sử dụng docker exec để chạy MySQL trực tiếp từ container
   docker exec -it mysqlgo8 mysql -uroot -proot1234

3. Docker tạo redis
   docker run --name redis -d -p 6381:6379 redis

4. Docker tạo kafka
   docker run --name kafka -d -p 19092:9092 kafka

5. Link swag
   http://localhost:8082/swagger/index.html

docker exec -it mysqlgo8 mysqldump -uroot -proot1234 --databases shopdevgo --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > migrations/shopdevgo.sql

goose -dir sql/schema create pre_go_crm_user_c sql

# Chạy Zookeeper

docker run --name zookeeper -d -p 2181:2181 -e ALLOW_ANONYMOUS_LOGIN=yes bitnami/zookeeper

# Chạy Kafka

docker run --name kafka -d -p 19092:9092 -e KAFKA_BROKER_ID=1 -e KAFKA_ZOOKEEPER_CONNECT=host.docker.internal:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:19092 -e ALLOW_PLAINTEXT_LISTENER=yes bitnami/kafka

# Chạy Kafka UI

docker run --name kafka-ui -d -p 9091:9091 -e KAFKA_CLUSTERS_0_NAME=local -e KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=host.docker.internal:19092 -e KAFKA_CLUSTERS_0_ZOOKEEPER=host.docker.internal:2181 provectuslabs/kafka-ui:latest

Loại Mô tả
:one Trả về một bản ghi duy nhất.
:many Trả về nhiều bản ghi.
:exec Thực thi câu lệnh, không quan tâm đến kết quả trả về.
:execresult: Thực thi truy vấn và trả về một số thông tin liên quan đến kết quả.
:execrows Thực thi câu lệnh và trả về số hàng bị ảnh hưởng.
:batchexec Thực thi nhiều câu lệnh trong một lô (batch).
:scalar Trả về một giá trị duy nhất (như COUNT, SUM).
:copyfrom Dùng với PostgreSQL để sao chép dữ liệu hàng loạt.

# chạy redis

docker exec -it redis redis-cli

# chạy tất cả container

docker start $(docker ps -aq)

# dừng tất cả container

docker stop $(docker ps -aq)

