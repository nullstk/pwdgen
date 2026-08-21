"""Minimal example for PwdGen."""

from pwdgen import pwdgen


def main():
 runner = pwdgen({"name": "PwdGen", "dry_run": False})
 result = runner.execute()
 print(result)


if __name__ == "__main__":
 main()