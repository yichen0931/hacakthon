
``` mermaid
flowchart LR

subgraph v0["Vendor"]
    v1["Open discount"]
    v2["Set discounted time"]
    v3["Set discounted price"]
    v4["Set discounted quantity"]
    v5["See all meal items"]
end

subgraph c0["Customer"]
	c1["Look through list of vendors"]
    c2["Set discounted time"]
    c3["Set discounted price"]
	c4["Set discounted quantity"]
	c5["See all meal items"]
end

    customer["Customer"]
    order["Order"]
	vendor["Vendor"]
	meal["Meal"]
	discount["Discount"]
    ordermeals["Ordermeals"]
    
	customer@{ shape: disk}
    order@{ shape: disk}
    vendor@{ shape: disk}
	meal@{shape: disk}
	discount@{shape: disk}
    ordermeals@{ shape: disk}


v1 --> p1["PUT"] --> vendor
v2 --> p2["PUT"] --> vendor
v3 --> p3["POST"] --> discount
v4 --> p4["PUT"] --> discount
v5 --> p5["GET"] --> meal

c1 --> p6["GET"] --> vendor
c2 --> p7["GET"] --> discount
c3 --> p8["POST"] --> order
c3 --> p9["POST"] --> ordermeals

```