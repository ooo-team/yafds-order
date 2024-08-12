-- public.customers definition

-- Drop table

DROP TABLE IF EXISTS public.orders;

CREATE TABLE public.orders (
	id int8 NOT NULL,
	customer_id int8 not null,
	restraunt_id int8 not null,
	courier_id int8 null,
	status int2 not null,
	timestamp timestamp not null
);
CREATE UNIQUE INDEX xpkorders ON public.orders USING btree (id);
CREATE INDEX xie1orders ON public.orders USING btree (customer_id);
CREATE INDEX xie2orders ON public.orders USING btree (restraunt_id);
CREATE INDEX xie3orders ON public.orders USING btree (courier_id);

