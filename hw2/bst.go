package main

import (
	"fmt"
)

type Node struct {
	inside int
	parent *Node
	right  *Node
	left   *Node
}

type bst struct {
	root *Node
}

func (t *bst) Add(el int) {
	came := &Node{inside: el}
	if t.root == nil {
		t.root = came
	} else {
		now := t.root // 4tobi ni4ego ne change
		// we need to check if inside variable is greater or lower that ours and than think where 2 put us
		// check here if right is empty if it is put our node else compare to nodes and do the same check
		// ну че тут вайл тру и брейкать го траить
		for {
			if el > now.inside {
				if now.right == nil { // is it empty? yeah, sit here.
					now.right = came
					came.parent = now.right
					break
				} else {
					now = now.right // no bro it's taken find a new place.
				}
			} else if el < now.inside {
				if now.left == nil {
					now.left = came
					came.parent = now.left
					break
				} else {
					now = now.left
				}
			}
		}
	}

}

func (t *bst) isExist(el int) bool {
	now := t.root // 4tobi ne menyat' nothing --> nujna peremennaya
	if now == nil {
		return false
	} else {
		if el == now.inside {
			return true
		} else {
			for {
				if el > now.inside {
					if now.right == nil {
						return false
					} else {
						now = now.right
					}
				} else if el < now.inside {
					if now.left == nil {
						return false
					} else {
						now = now.left
					}
				}
			}
		}
	}
}

//est' feeling 4to Delete eto o4ko nado swapat koren if it got leaves
// voobwe firstly nado find that mf i kak iskat' prob like we Added same stuff just go right left right left
//check if it got leaves --> than its a papa or mama else --> it's a baby leaf just delete it.
// papa or mama --> swap w maxi child teper 4o teper napisat nado eto kakto

func (t *bst) Delete(el int) {
	rn := t.root    //gde mi right now
	var mommy *Node // 4tobi know whose child is that. Who is your daddy? mem4ik
	var child *Node // Who is your baby?

	for rn != nil && rn.inside != el { // pervi step cikl poiska child'a
		mommy = rn //remembering w who our child came
		if el < rn.inside {
			rn = rn.left
		} else {
			rn = rn.right
		}
	}
	//na exit'e we know est li u nas child or no and all his data
	if rn == nil { //
		return // idk mojno li pustoi return es 4e prost udalu ili flag here
	}

	if rn.left == nil && rn.right == nil { // u child'a net friends
		if rn == t.root { // check if our child == tree
			t.root = nil // fully deleted tree nafik
		} else if mommy.left == rn { // nu toest zawli v mommy and checked if it's elder or younger child
			mommy.left = nil
		} else {
			mommy.right = nil
		}
		return
	}
	// okay teper' esli u childa kakoi to friend is missing to mi mojem swap child and friend and that's it
	// esli u nas tam est oba firneda to we gotta check to 4to higher pisal pro max child

	if rn.left == nil || rn.right == nil { // only 1 friend
		if rn.left != nil {
			child = rn.left
		} else {
			child = rn.right
		}

		if rn == t.root {
			t.root = child // changed root esli udalyaem glavnogo kid'a

		} else if mommy.left == rn { // this block i est swap mi znaem c kakoi storoni stoyal child u parent'a na ego mest placing friend
			mommy.left = child //
		} else {
			mommy.right = child
		}
		child.parent = mommy
		return
	}

	//oba frienda est -- >
	startP := rn     // want to kick
	mini := rn.right // minimal but on right

	for mini.left != nil { // iwem minimal in full right mini tree, vspomni primer when we delete the root
		startP = mini // meyaem uselok sprava na uzelki cprava but w/o left children
		mini = mini.left
	}

	// found node w minimal inside (to est' samii levi in the right mini tree)
	rn.inside = mini.inside // change, no ne deleted . тк указатели не изменяем,кароч засунули на место удаляемого рофлан а рофлан не удалили еще
	if startP.left == mini {
		startP.left = mini.right // swaping то что нашли и его листьями если они есть с правым тк ну мы взяли савмый левый
	} else {
		startP.right = mini.right
	}
}

func main() {
	tree := &bst{}

	tree.Add(18)
	tree.Add(8)
	fmt.Println(tree.isExist(18))
	fmt.Println(tree.isExist(6))
	tree.Delete(8)
	fmt.Println(tree.isExist(8))

}
