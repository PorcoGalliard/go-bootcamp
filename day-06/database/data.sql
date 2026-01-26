-- SOAL NOMOR 1
select c.category_id, c.category_name, count(p.product_id) as total_product
from oe.categories c 
join oe.products p 
	on c.category_id  = p.category_id 
group by 
	c.category_id,
	c.category_name

-- SOAL NOMOR 2
select s.supplier_id, s.company_name, count(p.product_id) as total_product
from oe.suppliers s 
join oe.products p 
	on s.supplier_id = p.supplier_id 
group by
	s.supplier_id,
	s.company_name

-- SOAL NOMOR 3
select s.supplier_id, s.company_name, count(p.product_id) as total_product, to_char(avg(p.unit_price), '999D99') as avg_unit_price
from oe.suppliers s 
join oe.products p 
	on s.supplier_id = p.supplier_id 
group by
	s.supplier_id,
	s.company_name


-- SOAL NOMOR 5
select c.customer_id, c.company_name, count(o.customer_id) as total_order
from oe.customers c 
join oe.orders o 
	on c.customer_id = o.customer_id 
group by 
	c.customer_id,
	c.company_name
order by
	c.company_name asc

-- SOAL NOMOR 6
select order_id, customer_id, order_date, required_date, shipped_date, shipped_date - order_date as delivery_time
from oe.orders o 
where shipped_date - order_date > 7

-- SOAL NOMOR 7
select p.product_id, p.product_name, sum(od.quantity) as total_qty
from oe.products p 
join oe.order_details od 
	on p.product_id = od.product_id
group by
	p.product_id,
	p.product_name
order by
	total_quantity desc

-- SOAL NOMOR 8
select c.category_id, c.category_name, sum(od.quantity) as total_qty_ordered
from oe.categories c 
join oe.products p 
	on c.category_id = p.category_id 
join oe.order_details od 
	on p.product_id = od.product_id 
group by
	c.category_id,
	c.category_name
order by
	total_qty_ordered desc

-- SOAL NOMOR 9
with ordered_qty_by_category as (
	select c.category_id, c.category_name, sum(od.quantity) as total_qty_ordered
	from oe.categories c 
	join oe.products p 
		on c.category_id = p.category_id 
	join oe.order_details od 
		on p.product_id = od.product_id 
	group by
		c.category_id,
		c.category_name
)

select category_id, category_name, total_qty_ordered
from ordered_qty_by_category
where
	total_qty_ordered = (select min(total_qty_ordered) from ordered_qty_by_category)
	or
	total_qty_ordered = (select max(total_qty_ordered) from ordered_qty_by_category)
order by total_qty_ordered desc

-- SOAL NOMOR 10
select s.shipper_id, s.company_name, p.product_id, p.product_name, sum(od.quantity) as total_qty_ordered
from oe.shippers s 
join oe.orders o 
	on s.shipper_id = o.ship_via
join oe.order_details od 
	on o.order_id = od.order_id
join oe.products p 
	on p.product_id = od.product_id
group by 
	s.shipper_id,
	s.company_name,
	p.product_id 
order by s.company_name, p.product_name 

-- SOAL NOMOR 11
WITH qty_per_shipper_product AS (
  SELECT
    s.shipper_id,
    s.company_name,
    p.product_id,
    p.product_name,
    SUM(od.quantity) AS total_qty_ordered
  FROM oe.shippers s
  JOIN oe.orders o
    ON s.shipper_id = o.ship_via
  JOIN oe.order_details od
    ON o.order_id = od.order_id
  JOIN oe.products p
    ON p.product_id = od.product_id
  GROUP BY
    s.shipper_id,
    s.company_name,
    p.product_id,
    p.product_name
)

SELECT *
FROM (
  SELECT DISTINCT ON (shipper_id)
    shipper_id,
    company_name,
    product_id,
    product_name,
    total_qty_ordered
  FROM qty_per_shipper_product
  ORDER BY
    shipper_id,
    total_qty_ordered DESC
) max_qty

UNION ALL

SELECT *
FROM (
  SELECT DISTINCT ON (shipper_id)
    shipper_id,
    company_name,
    product_id,
    product_name,
    total_qty_ordered
  FROM qty_per_shipper_product
  ORDER BY
    shipper_id,
    total_qty_ordered ASC
) min_qty

ORDER BY shipper_id, total_qty_ordered DESC;

