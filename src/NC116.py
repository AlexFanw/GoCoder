class Solution:
    def solve(self , nums ):
        if nums == '0': return 0
        if not nums:return 1
        nums = str(nums)
        dp = [0] * len(nums)
        dp[0] = 1
        for i in range(1, len(nums)):
            if nums[i] != '0': # 区别于lc，0不参与编码,比如“80” -> 0
                dp[i] = dp[i-1]
            if nums[i-1] == '1' or(nums[i-1] == '2' and nums[i] <= '6'):
                if i >= 2:
                    dp[i] += dp[i-2]
                else:
                    dp[i] += 1
        # print(dp)
        return dp[len(nums)-1]