Name: hydra-worker-sort-by-number
Version: 1
Release: 0
Summary: hydra-worker-sort-by-number
Source0: hydra-worker-sort-by-number-1.0.tar.gz
License: MIT
Group: custom
URL: https://github.com/innotech/hydra-worker-sort-by-number
BuildArch: x86_64
BuildRoot: %{_tmppath}/%{name}-buildroot
Requires: libzmq3
%description
Sort instances by number.
%prep
%setup -q
%build
%install
install -m 0755 -d $RPM_BUILD_ROOT/usr/local/hydra
install -m 0755 /hydra-worker-sort-by-number $RPM_BUILD_ROOT/usr/local/hydra/hydra-worker-sort-by-number

install -m 0755 -d $RPM_BUILD_ROOT/etc/init.d
install -m 0755 /hydra-worker-sort-by-number-init.d.sh $RPM_BUILD_ROOT/etc/init.d/hydra-worker-sort-by-number
%clean
rm -rf $RPM_BUILD_ROOT
%post
echo   When finished, you may want to run \"update-rc.d hydra-worker-sort-by-number defaults\"
%files
/usr/local/hydra/hydra-worker-sort-by-number
/etc/init.d/hydra-worker-sort-by-number
