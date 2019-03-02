class Solution:
    words = (
        ('Ninety', 90),
        ('Eighty', 80),
        ('Seventy', 70),
        ('Sixty', 60),
        ('Fifty', 50),
        ('Forty', 40),
        ('Thirty', 30),
        ('Twenty', 20),
        ('Nineteen', 19),
        ('Eighteen', 18),
        ('Seventeen', 17),
        ('Sixteen', 16),
        ('Fifteen', 15),
        ('Fourteen', 14),
        ('Thirteen', 13),
        ('Twelve', 12),
        ('Eleven', 11),
        ('Ten', 10),
        ('Nine', 9),
        ('Eight', 8),
        ('Seven', 7),
        ('Six', 6),
        ('Five', 5),
        ('Four', 4),
        ('Three', 3),
        ('Two', 2),
        ('One',1)
    )
    units = (
        ('Billion', 1000000000),
        ('Million', 1000000),
        ('Thousand', 1000),
        ('Hundred', 100)
    )
    def numberToWords(self, num: int) -> str:
        
        if num == 0:
            return "Zero"

        return " ".join(self.convert(num))
    
    def convert(self, num):
        
        for text, unit in self.units:
            if num >= unit:
                yield " ".join(self.convert(int(num/unit)))
                yield text
                num %= unit

        for text, value in self.words:
            if num >= value:
                yield text
                num -= value
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
            
            ret = Solution().numberToWords(num)

            out = (ret);
            print(out)
        except StopIteration:
            break

if __name__ == '__main__':
    main()
