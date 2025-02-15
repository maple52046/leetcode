class Solution:
    def minWindow(self, s: str, t: str) -> str:
        
        if not s or not t or len(s) < len(t):
            return ""
        
        characters = []
        for c in t:
            characters.append(c)

        start = 0
        end = 0
        mw = False
        for left in range(len(s) - len(characters) + 1):
            if s[left] in characters:
                print("======= new round: s[{}] =========".format(left))
                index = characters.index(s[left])
                sorted = self.selectSort(characters, index, -1)

                # find full string
                right = left
                while right < len(s) and sorted < (len(characters) - 1):
                    if right == left:
                        right += 1

                    if s[right] in characters[sorted+1:]:
                        index = characters[sorted+1:].index(s[right]) + sorted + 1
                        print("find {} at t with index: {} ({})".format(s[right], index, characters))
                        sorted = self.selectSort(characters, index, sorted)

                    if sorted >= (len(characters) - 1):
                        #print("sorted columns out of range")
                        #print("right: {}".format(right))
                        break

                    right += 1

                #print(characters)
                print("left: {}, right: {}, sorted: {}".format(left, right, sorted))
                if sorted == len(characters)-1:
                    mw = True
                    if (end - start) > (right - left) or (end - start) == 0:
                        start = left
                        end = right

        return s[start:end+1] if mw else ""

    def selectSort(self, characters, index, sorted):
        sorted += 1
        if index != sorted:
            print("swap t[{}] and t[{}]".format(index, sorted))
            characters[sorted], characters[index] = characters[index], characters[sorted]
        print("sorted characters: {}".format(sorted))
        return sorted

if __name__ == "__main__":
    solution = Solution()

    data = []
    #data.append(("ADOBECODEBANC", "ABC"))
    #data.append(("adobecodebanc", "abdbac"))
    #data.append(("ab", "a"))
    #data.append(("a", "a"))
    #data.append(("a", "b"))
    data.append(("qxsncfwvbslazxuoxnedkukropehlwfbwxqycntdfgghecvdqbhcwtukkauwzzzvystcfohmupvastekunmqiidtjxriyqdyiyapohekxblrurbpgphoykjhjarhtwfduhvkpzumahdxagmivtxvgiepjvxetehnmczddurgdwdecrmzklxqubgfzfjslqizvheadvghrlnvcbxpnuhjxpqywzrkrbjokqpolxxflkapnzeatmltdbrackkwlvmwlbmxbjcmvebieilzwyszckzgulcihpgsssrtdvhaaligvvfrwaqyksegdjqmywfsoyotuxtwieefbjwxjpxvhcemnwzntgfjetdatyydecjgofdzudxbfbqsxpfsvmebijcbhcemfnuvtzupcrthujbuyiovvswdbagjdkxkyrftqbktvdcpogloqajlsgquiyfljcxjveengogbulsitexjeixwqpszoxkzzkiuooiweqxydqjywiiaqhyhwrgkosloetktjuponposfxrdhzdyibhesprjjczoyjhhgyxtnmlulextdatnecsyrlhangonsxxywutligguldpqgiemkymdjufycumwtjupfpdowjkjozzwjdivbvymrdlvzzstkmkpenfcyplnqkjzquutrsgiytdxsvsckftquzstqdihnrgfnbbevjovcutupnyburrpsjijmsqclyjzzk", "fxtusxonrfdrhxjamdkwm"))

    for s, t in data:
        print(solution.minWindow(s, t))
        print("--------- Ans --------")