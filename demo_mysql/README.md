## 测试表结构
### 建表语句
```sql
CREATE TABLE water_meter(water_meter_id int primary key auto_increment, water_meter_name text, location VARCHAR(30) not null default 'UNKNOWN');
```