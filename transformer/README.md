# Local Transformer Libraries

This contains the various libraries used for the local transformer.

> **WARNING**: This code is still a work-in-progress, and highly experimental.
> There is no guarantee to its functionality. DO NOT USE

See
[this](https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-modifications.md)
for a high level description of the types of transformations that are done. Note
that not all have been implemented here yet!

## Packages

amp_packager
: This contains the compiled validator.proto from ampproject/amphtml. This is a
temporary solution for the time being. TODO(alin04): Add github issue.

layout
: Applies the AMP Layout algorithm.

printer
: Emits the DOM tree

transform
: Convenient entrypoint to the local transformer library.

transformer
: The individual transformers

