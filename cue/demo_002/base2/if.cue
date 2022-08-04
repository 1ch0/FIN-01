// 条件控制语句
// 没有 else，所有判断都会被执行
app: {
    name: string
    tech: string
    mem: int

    if tech == "react" {
        tier: "frontend"
    }
    if tech != "react" {
        tier: "backend"
    }

    if mem < 1Gi {
        footprint: "small"
    }
    if mem >= 1Gi && mem < 4Gi {
        footprint: "medium"
    }
    if mem  >= 4Gi {
        footprint: "large"
    }
}