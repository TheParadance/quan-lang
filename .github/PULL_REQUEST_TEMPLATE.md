# 📌 Pull Request Template – QuanLang Interpreter

---

## ✅ Summary

> Briefly describe what this PR does.

- What feature or bug does it address?
- What part(s) of the interpreter are affected?
- Is it a breaking change?

---

## 🧠 Motivation

> Why is this change necessary? What problem or feature does it address?

---

## 🔍 Changes Made

- [ ] New Expression Type(s)
  - [ ] `MemberExpr`
  - [ ] `AssignExpr` with extended target support (e.g., `a.x = 10`)
- [ ] Parser updates (e.g., `parsePrecedence`, `parsePrimary`)
- [ ] Evaluator logic for member assignments and expressions
- [ ] Object evaluation enhancements
- [ ] Tests added or updated

---

## 🧪 How to Test

```qlang
a = { x: 10 }
a.x = 20
print(a.x)  # should print 20
```

- Run interpreter with this script
- Confirm output is correct
- Confirm no crashes or regressions

---

## 🛠️ Checklist

- [ ] Code compiles and runs correctly
- [ ] All tests pass
- [ ] Code follows naming and structure conventions
- [ ] Docs updated (if needed)
- [ ] No debug prints or commented code in final version

---

## 📚 Related Issues / PRs

- Fixes #[issue_number]
- Related to #[other_pr_or_issue]

---

## 👀 Reviewer Notes

> Anything specific the reviewer should look at or keep in mind?

---
