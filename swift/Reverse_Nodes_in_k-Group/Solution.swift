/**
 * Definition for singly-linked list.
 */
public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init(_ val: Int) {
        self.val = val
        self.next = nil
    }
}

class Solution {
    func reverseKGroup(_ head: ListNode?, _ k: Int) -> ListNode? {

        var first:ListNode? = head      
        var newHead:ListNode?
        var current:ListNode? = head
        var previous:ListNode?
        var count = 0

        while current != nil {
            count += 1
            if count >= k {
                let (_head, _tail) = self.reverse(first, k)

                if newHead == nil {
                    newHead = _head
                }

                if let pv = previous {
                    pv.next = _head
                }
                previous = _tail

                _tail!.next = first!.next
                first = _tail!.next
                current = _tail
                count = 0
            }
            current = current!.next
        }

        if newHead == nil {
            newHead = head
        }

        return newHead
    }

    func reverse (_ head: ListNode?, _ k: Int) -> (ListNode?, ListNode?) {
        var first:ListNode? = head
        let current:ListNode? = head
        var next:ListNode? = head!.next

        for _ in 1..<k {
            current!.next = next!.next
            next!.next = first
            first = next
            next = current!.next
        }

        return (first, current)
    }
}

public class LinkedList {
	fileprivate var head: ListNode?
	private var tail: ListNode?

    public var first: ListNode? {
        return head
    }

	public func append(_ value: Int) {
        let newNode = ListNode(value)
        if let tailNode = tail {
            tailNode.next = newNode
        } else {
            head = newNode
        }
        tail = newNode

	}
}

func answer(_ result: ListNode?) {
    var res:ListNode? = result
    var ans:Array<Int> = []

    while (res != nil) {
        ans.append(res!.val)
        res = res!.next
    }

    print(ans)
}

func data(_ elments: Array<Int>) -> ListNode? {
    let data = LinkedList()
    for element in elments {
        data.append(element)
    }
    return data.first
}

func main() {
    let solution = Solution()
    var result:ListNode?

    // case 1
    result = solution.reverseKGroup(data([1,2,3,4,5]), 2)
    answer(result)

    // case 2
    result = solution.reverseKGroup(data([1]), 2)
    answer(result)

    // case 3
    result = solution.reverseKGroup(data([8,9,8,8,5,7,7,0,3,5]), 3)
    answer(result)



}

main()
