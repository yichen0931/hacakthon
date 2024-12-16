
``` mermaid
flowchart LR

subgraph v0["Vendor"]
    v1["Open discount"]
    v2["Set discounted time"]
    v3["Set discounted price"]
    v4["Set discounted quantity"]
    v5["See all meal items"]
    v6["See if discount is active"]
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


v1 --> vv1["PUT"] --> vendor
v2 --> vv2["PUT"] --> vendor
v3 --> vv3[" vvST"] --> discount
v4 --> vv4["PUT"] --> discount
v5 --> vv5["GET"] --> meal
v6 --> vv6["GET"] --> vendor

c1 --> cc1["GET"] --> vendor
c2 --> cc2["GET"] --> discount
c3 --> cc3["POST"] --> order
c3 --> cc4["POST"] --> ordermeals

```