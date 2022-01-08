drop table if exists produtos;

create table produtos(
	id serial primary key,
	nome varchar,
	descricao  varchar,
	preco	decimal,
	quantidade integer
);


insert into produtos (nome, descricao, preco, quantidade) values('Camiseta', 'Preta', 19, 3);
insert into produtos (nome, descricao, preco, quantidade) values('Fone', 'Totop', 129.9, 5);

select * from produtos;

  