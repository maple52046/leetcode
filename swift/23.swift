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
    func mergeKLists(_ lists: [ListNode?]) -> ListNode? {

        switch lists.count {
            case 0:
                return nil

            case 1:
                return lists[0]

            case 2:
                return self.merge(leftNode: lists[0], rightNode: lists[1])
            
            default:
                let m = lists.count/2
                let left = self.mergeKLists(Array(lists[0..<m]))
                let right = self.mergeKLists(Array(lists[m..<lists.count]))
                return self.merge(leftNode: left, rightNode: right)
        }
        
    }
    
    func merge(leftNode: ListNode?, rightNode: ListNode?) -> ListNode? {

        var head: ListNode?
        var tail: ListNode?
        var current: ListNode? 

        var left:ListNode? = leftNode
        var right:ListNode? = rightNode

        if left == nil {
            return right
        } else if right == nil {
            return left
        }

        while (left != nil && right != nil) {
            if left!.val <= right!.val {
                current = left
                left = left!.next
            } else {
                current = right
                right = right!.next
            }
            
            if let tailNode = tail {
                tailNode.next = current
            } else {
                head = current
            }

            tail = current
        }

        current!.next = (left != nil) ? left : right
        return head
    }
}

public class LinkList {
    fileprivate var head: ListNode?
    var tail: ListNode?

    public var first: ListNode? {
        return head
    }

    public var end: ListNode? {
        return tail
    }

    public var isEmpty: Bool {
        return head == nil
    }

    public func append(val: Int) {
        let newNode = ListNode(val)

        if let tailNode = tail {
            tailNode.next = newNode
        } else {
            head = newNode
        }
        tail = newNode
    }
}

func main() {

    // case 1
    let lists = [[1,4,5], [1,2,3], [2,6]]
    var nodes:[ListNode?] = []

    for l in lists {
        let node = LinkList()
        for val in l {
            node.append(val: val)
        }
        nodes.append(node.first)
    }

    let solution = Solution()
    let result: ListNode? = solution.mergeKLists(nodes)
    var res = result
    var answer:Array<Int> = []

    while (res != nil) {
        answer.append(res!.val)
        res = res!.next
    }
    print(answer)

    //case 2
    /* var nodes2:[ListNode?] = []
    nodes2.append(nil)
    nodes2.append(ListNode(1))
    let result2 = solution.mergeKLists(nodes2)
    var res2 = result2

    while (res2 != nil) {
        print("\(res2!.val) -> ")
        res2 = res2!.next
    } */

}

main()
