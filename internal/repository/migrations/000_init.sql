-- public.customers definition

-- Drop table

DROP TABLE IF EXISTS public.orders;

CREATE TABLE public.orders (
	id int8 NOT NULL,
	customer_id int8 not null,
	restaurant_id int8 not null,
	courier_id int8 null,
	status int2 not null,
	timestamp timestamp not null
);
CREATE UNIQUE INDEX xpkorders ON public.orders USING btree (id);
CREATE INDEX xie1orders ON public.orders USING btree (customer_id);
CREATE INDEX xie2orders ON public.orders USING btree (restaurant_id);
CREATE INDEX xie3orders ON public.orders USING btree (courier_id);


DROP TABLE IF EXISTS public.order_states;

CREATE TABLE public.order_states (
	id int2 NOT NULL,
	descriprion varchar
);
CREATE UNIQUE INDEX xpkorderstates ON public.orders USING btree (id);
