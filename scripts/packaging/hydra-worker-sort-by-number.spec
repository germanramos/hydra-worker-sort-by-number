%define debug_package %{nil}
Name: hydra-worker-sort-by-number
Version: 2.0.0
Release: 1
Summary: hydra-worker-sort-by-number
Source0: hydra-worker-sort-by-number-2.0.0.tar.gz
License: MIT
Group: custom
URL: https://github.com/innotech/hydra-worker-sort-by-number
BuildArch: x86_64
BuildRoot: %{_tmppath}/%{name}-buildroot
%description
Sort instances by number.
%prep
%setup -q
%build
%install
install -m 0755 -d $RPM_BUILD_ROOT/usr/local/hydra
install -m 0755 hydra-worker-sort-by-number $RPM_BUILD_ROOT/usr/local/hydra/hydra-worker-sort-by-number

install -m 0755 -d $RPM_BUILD_ROOT/etc/init.d
install -m 0755 hydra-worker-sort-by-number-init.d.sh $RPM_BUILD_ROOT/etc/init.d/hydra-worker-sort-by-number

install -m 0755 -d $RPM_BUILD_ROOT/etc/hydra
install -m 0644 hydra-worker-sort-by-number.conf $RPM_BUILD_ROOT/etc/hydra/hydra-worker-sort-by-number.conf
%clean
rm -rf $RPM_BUILD_ROOT
%post
echo   You should edit config file /etc/hydra/hydra-worker-sort-by-number.conf
echo   When finished, you may want to run \"update-rc.d hydra-worker-sort-by-number defaults\"
%files
/usr/local/hydra/hydra-worker-sort-by-number
/etc/init.d/hydra-worker-sort-by-number
%dir /etc/hydra
/etc/hydra/hydra-worker-sort-by-number.conf
/etc/init.d/hydra-worker-sort-by-number
