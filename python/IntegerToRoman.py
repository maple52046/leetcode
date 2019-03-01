class Solution:
    def intToRoman(self, num: int) -> str:
        
        return ''.join(self.convert(num))
        
    def convert(self, num: int) -> str:
        roman = {
            1: 'I',
            4: 'IV',
            5: 'V',
            9: 'IX',
            10: 'X',
            40: 'XL',
            50: 'L',
            90: 'XC',
            100: 'C',
            400: 'CD',
            500: 'D',
            900: 'CM',
            1000: 'M'
        }
        
        thresholds = list(roman.keys())
        
        index = -1
        while num > 0:
            threshold = thresholds[index]
            if num >= threshold:
                yield roman[threshold]
                num -= threshold
            else:
                index -= 1

def main():
    import sys
    import io
    def readlines():
        for line in io.TextIOWrapper(sys.stdin.buffer, encoding='utf-8'):
            yield line.strip('\n')

    lines = readlines()
    while True:
        try:
            line = next(lines)
            num = int(line);
            
            ret = Solution().intToRoman(num)

            out = (ret);
            print(out)
        except StopIteration:
            break

if __name__ == '__main__':
    main()
