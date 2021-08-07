package Swap_Adjacent_in_LR_String

func canTransform(start string, end string) bool {
    if len(start) != len(end) {
        return false
    }

    if len(start) == 0 || len(end) == 0{
        return true
    }

    length := len(start)
    getNext := func(str string,index int)(byte,int){
        for index < len(str) && str[index] == 'X' {
            index++
        }

        if index >= len(str) {
            return 'X',index
        }

        return str[index],index+1
    }

    p1,p2 := 0,0
    for p1 < length || p2 < length {
        var b1,b2 byte
        b1,p1 = getNext(start,p1)
        b2,p2 = getNext(end,p2)

        if b1 != b2 {
            return false
        }

        if b1 == 'R' && p1>p2 {
            return false
        }

        if b1 == 'L' && p1 < p2 {
            return false
        }

    }

    return true

}

