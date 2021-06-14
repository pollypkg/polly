# Contributing

Polly uses GitHub for [discussions](https://github.com/pollypkg/polly/discussions) and pull requests. In this early phase, there are a few different avenues looking for contributions in the following areas:

1. [Writing test cases/examples](#Exploration-by-example)
2. [Spec development](#Specification-development)
3. [Adding an object type](#Adding-an-object-type)

And of course, there's also the [`good first issue`](https://github.com/pollypkg/polly/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22) and [`help wanted`](https://github.com/pollypkg/polly/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22) issue labels!

## Exploration by example

Polly is in an exploratory prototype phase. We're still deciding how some of its basic mechanisms will work. The fastest way for answers to reveal themselves is by actually trying to make polly packages (pops), and see what issues arise.

This is also a great way for (perseverent) folks to get a sense of what Polly is: try to make a pop, and put it up in a PR for discussion of questions or insights you may have had. (_Especially_ if the result really didn't work out!)

If you want to give this a shot, here's a procedure:

1. [Install the `cue` CLI tool](https://cuelang.org/docs/install/).
2. Clone this repository.
3. Pick a system to observe. (Strongly recommended: pick a system with an [existing](https://monitoring.mixins.dev/) [mixin](https://github.com/grafana/jsonnet-libs) to migrate)
4. Pick some generally sensible name for the pop, based on the system you intend to observe. Create a directory with that name under [`examples/`](https://github.com/pollypkg/polly/tree/main/examples).
5. Look at other examples to see how the pop is structured and how to use the various object types.
6. Create your own! If you're adapting an existing mixin, it's easiest to pick out some subset of the objects to copy over. Consider taking notes about what you're thinking as you go through.
7. [Send a PR with your pop, using this template!]() It's got questions to help guide your response. (It may be worth reading the questions beforehand to help guide your note-taking.)

Example pops may eventually mature to the point where they're moved out into their own repositories/wherever makes sense. For now, this keeps all the attention in one place, which is good for maximizing learning!

## Specification development

The canonical polly specification - expressed in CUE - [resides](https://github.com/pollypkg/polly/tree/main/schema) in this repository. All changes to the specification are conducted through reviews and pull requests against this repository.

The specification is being actively developed (read: unstable). While that is the case, we expect to proceed with a relatively informal pull request workflow: think of a change, make a PR, discuss on the fly. As the specification matures, we expect to put in place a more formal proposal process, especially for more significant proposals.

## Adding an object type

Polly is designed to be extensible in specific ways. Chief among these is the ability to add new types of configuration objects to the specification.

This is less of a focus for now, in Polly's early phases of development. However, anyone interested in adding configuration for an appropriate type of object is welcome to make a proposal - we'll treat it as a forcing function to clarify the rules! For now, we're pretty sure this is key:

1. The object's schema must be defined using scuemata
2. The mechanism for (integrating a signal)[https://github.com/pollypkg/polly/discussions/6] into the config object must be defined