#class Solution:
#    def findRotateSteps(self, ring, key):
#        """
#        :type ring: str
#        :type key: str
#        :rtype: int
#        """
#        current = 0
#        length = len(ring)
#        count = len(key)
#        
#        for k in key:
#
#            print(current, k)
#            if k != ring[current]:
#                steps = 1
#                while True:
#                    right = (current+steps)%length
#                    left = (current-steps)%length
#                    
#                    if k in [ring[right], ring[left]]:
#                        break
#
#                    steps += 1
#                
#                count += steps
#        
#        return count

class Solution:
    def findRotateSteps(self, ring, key):
        """
        :type ring: str
        :type key: str
        :rtype: int
        """
        count = 0
        length = len(ring)
        distance_map = {}
        # if key[0] != ring[current]:
        #     steps = 1
        #     while True:
        #         right = (current+steps)%length
        #         left = (current-steps)%length
        #         if key[0] == ring[right]:
        #             current = right
        #             break
        #         elif key[0] == ring[left]:
        #             current = left
        #             break
        #         else:
        #             steps += 1
        #     count += steps
        # print("count: {}, current index: {}".format(count, current))

        distance_map = [{}] * length
        for i in range(length):
            for j in range(length):
                if i != j:
                    right = (i+j)%length
                    left = (i-j)%length

        
        for i in range(length):
            print("\nring[{}] => {}".format(i, ring[i]))
            steps = 1
            shift = [0, 0]
            if ring[i] not in distance_map:
                distance_map[ring[i]] = {}
            while True:
                right = (i+steps)%length
                left = (i-steps)%length
                print("right: {}, left: {}".format(right, left))

                if ring[i] != ring[right]:
                    distance = steps + shift[0]
                    print("distance of {} -> {}: {}".format(ring[i], ring[right], distance))
                    previous = distance_map[ring[i]].setdefault(ring[right], distance)
                    if distance < previous:
                        distance_map[ring[i]][ring[right]] = distance
                    shift[0] = 0
                else:
                    print("right shift")
                    shift[0] += 1

                if ring[i] != ring[left]:
                    distance = steps + shift[1]
                    print("distance of {} -> {}: {}".format(ring[i], ring[left], distance))
                    previous = distance_map[ring[i]].setdefault(ring[left], distance)
                    if distance < previous:
                        distance_map[ring[i]][ring[left]] = distance
                    shift[1] = 0
                else:
                    print("right shift")
                    shift[1] += 1

                if abs(right - left) <= 1 or (abs(right - left)+1) >= length:
                    print("break because right: {}, left: {}".format(right, left))
                    break
                steps += 1
                #input("continue...")
        print(distance_map)


        current = ring[0]
        for k in key:
            if k != current:
                print("trun {} time to {}".format(distance_map[current][k], k))
                count += distance_map[current][k]
                current = k
            print("press button")
            count += 1

        return count

if __name__ == "__main__":
    s = Solution()
    #print(s.findRotateSteps("godding", "gd"))
    #print(s.findRotateSteps("abcde", "ade"))
    #print(s.findRotateSteps("edcba", "abcde"))
    print(s.findRotateSteps("godding", "godding"))
    #print(s.findRotateSteps("nyngl", "yyynnnnnnlllggg"))