# Tag percentage Timewarrior extension

This [Timewarrior extension](https://timewarrior.net/docs/extensions/) calculates the percentage of time registered for a particular tag on the total for each day in the input intervals.

## Installation

1. Download the latest executable for your operating system from
   the [releases page](https://github.com/crossbone-magister/tag-percentage/releases).
2. Add it to the Timewarrior extension folder as described in the [documentation](https://timewarrior.net/docs/api/).
3. Verify that the extension is active and installed by running `timew extensions`.

## Configuration

Add the entry `reports.tagpercentage.target` Timewarrior configuration to allow this extension to work.
Its value represents the tag to look for when calculating percentages.

For example:
```shell
timew config reports.tagpercentage.target project1
```

Sets the label `project1` as the target for this extension.


## Usage

In a terminal window, run `timew tag-percentage`. An example output could be:

```bash
2024-01-01 - project1: 80.39%
2024-01-02 - project1: 71.12%
2024-01-03 - project1: 23.67%
2024-01-04 - project1: 45.96%
2024-01-05 - project1: 85.71%
2024-01-06 - project1: 32.56%
Average percentage: 56.57%
```
