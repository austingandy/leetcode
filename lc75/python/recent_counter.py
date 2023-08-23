class RecentCounter:
    
    def __init__(self):
        self.stack = []
    
    def ping(self, t: int) -> int:
        i = 0
        self.stack.append(t)
        while i < len(self.stack) and self.stack[i] < t-3000:
            i += 1
        self.stack = self.stack[i:]
        return len(self.stack)


obj = RecentCounter()
for t in [1, 100, 3001, 3002]:
    print(obj.ping(t))
param_1 = obj.ping(t)
