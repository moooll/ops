CREATE MATERIALIZED VIEW order_products_info AS
  SELECT
    p.id as product_id,
    p.name as product_name,
    p.description as product_description,
    op.order_id as order_id,
    op.quantity as quantity
  FROM order_products as op
    LEFT JOIN products p ON p.id = op.product_id;

CREATE OR REPLACE FUNCTION refresh_order_products()
  RETURNS TRIGGER LANGUAGE PLPGSQL
  AS $$
  BEGIN
    REFRESH MATERIALIZED VIEW order_products_info;
    RETURN NULL;
  END $$;

CREATE TRIGGER refresh_order_products
  AFTER INSERT OR UPDATE OR DELETE OR TRUNCATE
  ON order_products FOR EACH STATEMENT
  EXECUTE PROCEDURE refresh_order_products();

CREATE TRIGGER refresh_order_products
  AFTER INSERT OR UPDATE OR DELETE OR TRUNCATE
  ON orders FOR EACH STATEMENT
  EXECUTE PROCEDURE refresh_order_products();

CREATE TRIGGER refresh_order_products
  AFTER INSERT OR UPDATE OR DELETE OR TRUNCATE
  ON products FOR EACH STATEMENT
  EXECUTE PROCEDURE refresh_order_products();
