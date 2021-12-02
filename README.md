# dive-arch
Dive archive - as in diving from a board, not with SCUBA gear.

# Purpose
We want to archive "all" competition dives a diver makes, so that the diver can go back in history
to watch their old dives, perhaps see how they improve etc. We would like to integrate this with
some of the systems out there that log dive competitions. scores etc, so that we can also get this
information with the dives we have

# Future
It would be ideal if we could have a mobile phone application where the videographer could select
the current competition, the current diver, and then shoot video that is automatically archived
under the correct competition with the relevant metadata

# Security
Since dive videos and other things must be considered personal information, we must have a proper
login scheme for accessing these. The intention is to use OAuth JWT tokens, using, for example,
Auth0 as an auth provider
