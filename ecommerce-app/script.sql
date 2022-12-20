IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[users]') AND type in (N'U'))
DROP TABLE [dbo].[users]
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[products]') AND type in (N'U'))
DROP TABLE [dbo].[products]
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[orders]') AND type in (N'U'))
DROP TABLE [dbo].[products]
GO


create table users(
    ID int not null primary key IDENTITY(1,1),
    name varchar(100),
    email varchar(100),
    password varchar(100),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
)

create table products(
    ID int  not null primary key IDENTITY(1,1),
     name varchar(100),
    quantity int,
    description varchar(1000),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
)
create table orders(
    ID int not null primary key IDENTITY(1,1),
    userid int,
    productid int,
    quantity int,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
)