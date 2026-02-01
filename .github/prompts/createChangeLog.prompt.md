---
name: createChangeLog
description: Works through the release PR to compile a change long
argument-hint: versionNumber
---
# Instructions
- Follow EVERY step
- Do NOT start the following step until the previous step has finished
- Do NOT skip any steps
- This project is `Murex`, found on Github at https://github.com/lmorg/murex
- changelogs are found in `gen/changelog/`
- changelogs consist of 2 (two) files with the naming convention:
    1. `${VERSION}_doc.yaml`
    2. `${VERSION}.inc.md`
- `${VERSION}` is a version number and should follow the format: `v${MAJOR}.{MINOR}` where:
    1. `${MAJOR}` is a numeric value
    2. `${MINOR}` is a numeric value


# Step 1: Find Github Pull Request (PR)
- Check for open PRs in Github for only the Murex repository. We are only interested in a PR that is named like `v7.1`. Ignore all other PRs.
- There should only be one PR. If there is more than one PRs, or you are unsure which PR to check, then stop all steps and ask which PR to read.
- Do not read multiple PRs
- Do not read other repositories other than `murex`

# Step 2: Read all git commits, then summarize changes
From the selected PR:
- Read every git commit and create a list of changes.
- Use git commit messages as a clue for the change.
- If a commit message has a value like `#123` (`#` followed by a number) then this is a Github Issue. You should also read that Github issue to understand the change.
- Multiple git commits might be part of the same change.
- A git commit might contain multiple changes.

You should output a bullet point list of each change, with a short summary in the following format:
```
- ${COMPONENT}: ${SUMMARY} ([issue](${ISSUE}))
```
    1. `${COMPONENT}` is a one word title for the system that is being changed
    2. `${SUMMARY}` is the summary you generate
    3. `${ISSUE}` is a URL to the Github issue -- if an change has a related Github Issue

Some issues might be links to Github discussions.

Also include a reference to any contributors either creating related Github issues, github discussions, commits or pull requests. These users should be related to this version and who are not `@lmorg`.

# Step 3: Grouping changes

The changes should be grouped into the following groups:
1. **Breaking changes**: these are changes that might impact the users upgrading from previous versions of this project
2. **Features**: these are new features added to the project in this version
3. **Bug Fixes**: these are changes that do not introduce new features but that fix bugs or issues with existing features

# Step 5: Update changelog
This step refers only to `${VERSION}.inc.md`:
- This file should be in `gen/changelog/`.
- If a file does not exist, then you can create one.
- If a file does exist then review the contents of it and add any changes you've identified if they don't already exist in this document
- This is a markdown document.
- It's contents should match the following:
```
## Breaking Changes

- ${COMPONENT}: ${SUMMARY} ([issue](${ISSUE}))

## v${VERSION}.xxxx

### Features

- ${COMPONENT}: ${SUMMARY} ([issue](${ISSUE}))

### Bug Fixes

- ${COMPONENT}: ${SUMMARY} ([issue](${ISSUE}))

## Special Thanks

Thank yous for this release go to ${LIST_OF_USERS_WHO_CONTRIBUTED} for your code, testing and feedback.

Also thank you to everyone in the [discussions group](https://github.com/lmorg/murex/discussions) and all who raise bug reports.

You rock!
```

# Step 6: Update summary

This step refers only to `${VERSION}_doc.yaml`:
- This file should be in `gen/changelog/`.
- If a file does not exist, then you can create one.
- If a file does exist then review the contents of it and add any changes if you think it alters the nature of the version summary.
- This is a YAML document.
- It's contents should match the following:

```
- DocumentID: ${VERSION}
  Title: >-
    ${VERSION}
  CategoryID: changelog
  DateTime: ${VERSION}
  Summary: >-
    ${VERSION_SUMMARY}
  Description: |-
    {{ include "gen/changelog/${VERSION}.inc.md" }}
  Related:
    - CONTRIBUTING
``` 
- ${VERSION} should be replaced with the version number.
- ${DATE} should be replaced with the current date and time, formatted like the following example: 2025-10-23 22:35
- ${VERSION_SUMMARY} will be a summary of all the changes in this version.

