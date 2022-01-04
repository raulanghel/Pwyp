print("test")

for i in range(5):
    print(i * 3)

x = ["a", "b", "c"]
y = [1, 2, 3]
z = x + y

print(len(z), z)


class Mouse:
    def __init__(self) -> None:
        self.feet = 4
