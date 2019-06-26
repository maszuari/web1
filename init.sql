CREATE DATABASE testwebdb WITH ENCODING 'UTF8' LC_COLLATE='en_US.utf8' LC_CTYPE='en_US.utf8'; 
CREATE USER web1 with encrypted password 'secret';
GRANT ALL PRIVILEGES ON DATABASE testwebdb to web1;
