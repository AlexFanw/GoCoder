package main

func main() {
	
}
func isMatchSub(rep string, strb string, match [][]bool){
	match[0][0] = true

	for i:=1;i<=len(rep);i++{
		if rep[i-1] == '*'{
			match[0][i] = match[0][i-1]
		}else{
			match[0][i] = false
		}
	}
	for i:=1;i<=len(strb);i++{
		for j:=1;j<=len(rep);j++{
			if rep[j-1] == '?' || rep[j-1] == strb[i-1]{
				match[i][j] = match[i-1][j-1]
			}else if rep[j-1] == '*'{
				m1 := match[i-1][j]
				m2 := match[i-1][j-1]
				m3 := match[i][j-1]
				match[i][j] = m1 || m2 || m3
			}else{
				match[i][j]=false
			}
		}
	}
}


func isMatch( s string ,  p string ) bool {
	// write code here
	if len(s) == 0 && len(p) ==0 {
		return true
	}
	if len(s) == 0{
		for i:=0;i<len(p);i++{
			if p[i] != '*'{
				return false
			}
		}
		return true
	}
	if len(p) == 0{
		return false
	}

	mtchList := make([][]bool, len(s)+1)
	for i:=0;i<len(mtchList);i++{
		mtchList[i] = make([]bool, len(p)+1)
	}
	isMatchSub(p, s, mtchList)
	return mtchList[len(s)][len(p)]
}
