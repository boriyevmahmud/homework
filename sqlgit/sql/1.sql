create table customers(customer_id uuid DEFAULT uuid_generate_v4 (),
firstName varchar(30),lastName varchar(30),userName varchar(30),email varchar(30),
gender varchar(15),birthdate date,password varchar(20),
status varchar(50),PRIMARY KEY(customer_id));


create table adress (adress_id uuid DEFAULT uuid_generate_v4 (),customer_id uuid,
country varchar(20),city varchar(20),district varchar(30),postalCodes int,
PRIMARY KEY(adress_id),
   CONSTRAINT fk_customer
      FOREIGN KEY(customer_id) 
    REFERENCES customers(customer_id)
    ON DELETE CASCADE 
);


create table product (product_id uuid DEFAULT uuid_generate_v4 (),customer_id uuid,
name varchar(20),cost int,
orderNumber int,amount int, currency varchar(20),rating int,
PRIMARY KEY(product_id),
   CONSTRAINT fk_customer
      FOREIGN KEY(customer_id) 
    REFERENCES customers(customer_id)
    ON DELETE CASCADE 
);

create table phone (phone_id uuid DEFAULT uuid_generate_v4 (),customer_id uuid,
 numbers int,code varchar(15),
 PRIMARY KEY(phone_id),
   CONSTRAINT fk_customer
      FOREIGN KEY(customer_id) 
    REFERENCES customers(customer_id)
    ON DELETE CASCADE 
);

create table typee()
select c.customer_id,firstname,lastname,username,country,city,district,postalcodes,
numbers,code,p.name,cost,ordernumber,amount,currency,rating,email,gender,birthdate,
password,status from customers c inner join adress a on c.customer_id=a.customer_id inner join 
product p on a.customer_id=p.customer_id inner join phone h on p.customer_id=h.customer_id inner join 
typee t on p.product_id=t.product_id  ;
