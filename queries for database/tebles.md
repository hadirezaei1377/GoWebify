table name is people :
                     columns:
                     id
                     first_name 
                     last_name

table name is emails :
                    columns:
                    id
                    people_id
                    email_address


commands in editor:

insert into people (first_name, last_name) values('John', 'Smith');
insert into people (first_name, last_name) values('Marry', 'Jones');

insert into emails (people_id, email_address) values('1', 'john@gmail.com');
insert into emails (people_id, email_address) values('2', 'marry@gmail.com');

