<h1 align='center'> blob-base </h1>

### Description
Service for storing blobs (custom user JSONs)  

### How to start
Firstly, export "KV_VIPER_FILE=./config.yaml", that is path to config   
Then, run `docker-compose -f docker-compose.yml up` to start database container  
After, run app wit `migrate up` to create necessary db  
Finally, start app with `run service`

### Add blob 
`POST /blobs/`  

![изображение](https://user-images.githubusercontent.com/108219165/188287743-8b368573-22e5-470a-8732-e0d97225df03.png)

### Get blobs 
`GET /blobs/`  

![изображение](https://user-images.githubusercontent.com/108219165/188287792-854ec724-8520-4b73-8352-878a54d341fe.png)

### Get blob by id 
`GET /blobs/{id}/`  

![изображение](https://user-images.githubusercontent.com/108219165/188287822-dd9fd1e4-d6ff-4b15-891b-1a25ec90d726.png)


### Delete blob by id 
`DELETE /bblobs/{id}/`  

![изображение](https://user-images.githubusercontent.com/108219165/188287840-4b28e1a1-42a8-4ced-aeae-1dc203403d9d.png)   


### Post in core  
`POST /accounts/`   

![изображение](https://user-images.githubusercontent.com/108219165/189493984-8741351f-38d7-4971-9c36-dbe4b1f4bc4c.png)
