#!/usr/bin/env python

import logging

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)

class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        
        alphbats = {}
        max_len = 0
        head = 0
        length = 0
        
        for i in range(len(s)):
            # By using setdefualt function,
            # "j" will equal to "i" when alphbats dict didn't have key "s[i]",
            # or "j" will be the original value of "alphbat[ s[i] ]".
            j = alphbats.setdefault(s[i], i)
            
            if j < head:
                # In this case, means the record index of target alphabt should be duprecated.
                # So, we updated index with new value.
                logger.debug("head = {}, so ignore original index {}: {}".format(head, s[i], alphbats[s[i]]))
                alphbats[s[i]] = i
                logger.debug("update index of {} to {}".format(s[i], i))
                
            if i != alphbats[s[i]]:
                # When "i" isn't equal to "j", means we found duplicate alphbats.
                
                # Update max length value
                length = len(s[head:i])
                if max_len < length:
                    max_len = length
                    logger.debug("update max length to {}".format(max_len))

                # Setup new head index
                if alphbats[s[i]] == head:
                    # In this case, means the first of duplicate alphbat is the head of fragment.
                    # So we offset the head index to next alphbat.
                    head += 1

                else:
                    # In this case, we need to move the head index to the next of first duplicate alphbat.
                    head = alphbats[s[i]] + 1

                logger.debug("update head to {}".format(head))
                alphbats[s[i]] = i
                    
            length = len(s[head:i+1])
            logger.debug("length: {} , string: {}".format(length, s[head:i+1]))
                
        if max_len < length:
            return length
        
        return max_len

if __name__ == "__main__":
    import sys

    formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    ch = logging.StreamHandler()
    ch.setLevel(logging.DEBUG)
    ch.setFormatter(formatter)
    logger.addHandler(ch)

    s = raw_input("Input string: ") if len(sys.argv) < 2 else sys.argv[1]
    obj = Solution()
    print "Result: {}".format(obj.lengthOfLongestSubstring(s))

# vim: ts=4 sw=4 expandtab
