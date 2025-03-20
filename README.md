### AstraLend-backend 项目说明

#### 此项目为astraLend 借贷项目的后端支持


gentool -dsn "root:123456@tcp(127.0.0.1:3306)/astra_lend" -tables "poolbases,pooldata" -outPath "D:\github projects\AstraLend-backend\api\models"

gormt -g -conn "root:123456@tcp(127.0.0.1:3306)/astra_lend?charset=utf8mb4&parseTime=True&loc=Local" -tables "poolbases,pooldata"  -outPath "/absolute/path/to/project/internal/repository/models"
