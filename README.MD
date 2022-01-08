# Test GORM DB Golang  

Aplicação para testar a biblioteca GORM para acesso a banco de dados em Golang.  


## Antes de Executar

É necessário subir um container com o banco Postgres.

``docker run --name db_container_name \  ``  
``    -e POSTGRES_USER=db_user \  ``  
``   -e POSTGRES_PASSWORD=db_pass \  ``  
``    -e POSTGRES_DB=db_name \  ``  
``    -p 5432:5432 \  ``  
``    -d postgres:14.1``  