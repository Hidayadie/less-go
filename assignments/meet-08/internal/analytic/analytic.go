// Package analytics menyediakan endpoint analisis (read-only) untuk
// pelanggan dan produk dari database Northwind.
package analytic

// Pendapatan = unit_price * quantity * (1 - discount), dihitung dari order_details.
// Kolom bertipe real di-cast ke numeric agar ROUND(…, 2) bersih dari artefak float.
// Setiap kolom WAJIB memakai alias: alias menjadi key JSON lewat json_agg.

// customerAnalysis: satu baris per pelanggan, termasuk yang belum pernah order
// (LEFT JOIN). COUNT(DISTINCT order_id) mencegah order terhitung ganda akibat
// fan-out join ke order_details; NULLIF mencegah pembagian dengan nol.
const customerAnalysis = `
SELECT c.customer_id,
       c.company_name,
       c.country,
       COUNT(DISTINCT o.order_id) AS total_orders,
       COALESCE(ROUND(SUM(od.unit_price * od.quantity * (1 - od.discount))::numeric, 2), 0) AS total_revenue,
       COALESCE(ROUND((SUM(od.unit_price * od.quantity * (1 - od.discount))
                       / NULLIF(COUNT(DISTINCT o.order_id), 0))::numeric, 2), 0) AS avg_order_value,
       MAX(o.order_date) AS last_order_date
FROM customers c
LEFT JOIN orders o ON o.customer_id = c.customer_id
LEFT JOIN order_details od ON od.order_id = o.order_id
GROUP BY c.customer_id, c.company_name, c.country
ORDER BY total_revenue DESC, c.company_name`

// productAnalysis: satu baris per produk, termasuk yang belum pernah terjual.
// needs_reorder hanya true untuk produk aktif yang stoknya sudah di bawah
// atau sama dengan reorder_level.
const productAnalysis = `
SELECT p.product_id,
       p.product_name,
       cat.category_name,
       COALESCE(SUM(od.quantity), 0) AS units_sold,
       COUNT(DISTINCT od.order_id) AS order_count,
       COALESCE(ROUND(SUM(od.unit_price * od.quantity * (1 - od.discount))::numeric, 2), 0) AS total_revenue,
       p.units_in_stock,
       p.reorder_level,
       (p.discontinued = 0 AND p.units_in_stock <= p.reorder_level) AS needs_reorder,
       p.discontinued
FROM products p
JOIN categories cat ON cat.category_id = p.category_id
LEFT JOIN order_details od ON od.product_id = p.product_id
GROUP BY p.product_id, p.product_name, cat.category_name,
         p.units_in_stock, p.reorder_level, p.discontinued
ORDER BY total_revenue DESC, p.product_name`
